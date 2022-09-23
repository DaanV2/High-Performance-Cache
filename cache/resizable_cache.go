package cache

import (
	"time"

	"go.uber.org/atomic"
)

// ResizableCache is a cache that can be resized
type ResizableCache[T KeyedObject] struct {
	// currentCache is the current cache
	currentCache CacheCleanable[T]
	// oldCache is the old cache
	oldCache CacheCleanable[T]
	// isResizing is a flag that indicates if the cache is currently resizing
	isResizing *atomic.Bool
	// settings is the settings of the cache
	settings ResizableCacheSettings
}

// NewResizableCache creates a new resizable cache
func NewResizableCache[T KeyedObject](settings ResizableCacheSettings) *ResizableCache[T] {
	result := &ResizableCache[T]{
		settings:   settings,
		isResizing: atomic.NewBool(false),
	}

	result.currentCache = NewFixedCache[T](settings.FixedCacheSettings)

	return result
}

// Set sets the item in the cache.
func (rc *ResizableCache[T]) Set(item T) bool {
	c := rc.currentCache

	if c != nil {
		return c.Set(item)
	}

	c = rc.oldCache
	if c != nil {
		return c.Set(item)
	}

	return false
}

// Get returns the item from the cache.
func (rc *ResizableCache[T]) Get(key string) (CacheItem[T], bool) {
	c := rc.currentCache

	if c != nil {
		return c.Get(key)
	}

	c = rc.oldCache
	if c != nil {
		return c.Get(key)
	}

	var r CacheItem[T]
	return r, false
}

// Gets the item from the cache or sets it if it does not exist.
func (rc *ResizableCache[T]) GetOrSet(key string, createFn func(key string) (T, error)) (CacheItem[T], error) {
	old := rc.oldCache
	current := rc.currentCache

	//If old cache is nil, we only need to check the current
	if old == nil {
		return current.GetOrSet(key, createFn)
	}

	if item, ok := current.Get(key); ok {
		return item, nil
	}
	if item, ok := old.Get(key); ok {
		return item, nil
	}

	return current.GetOrSet(key, createFn)
}

// Delete deletes the item from the cache.
func (rc *ResizableCache[T]) Delete(key string) bool {
	//No need to check oldCache, as it will be deleted in whole
	return rc.currentCache.Delete(key)
}

// Clear clears the cache.
func (rc *ResizableCache[T]) Clear() error {
	return rc.currentCache.Clear()
}

// Dispose disposes the object
func (rc *ResizableCache[T]) Dispose() {
	rc.currentCache.Dispose()

	if rc != nil {
		rc.oldCache.Dispose()
	}
}

// IsResizing returns true if the cache is currently resizing
func (rc *ResizableCache[T]) IsResizing() bool {
	return rc.isResizing.Load()
}

// Close closes the cache.
func (rc *ResizableCache[T]) ForEach(callback func(value CacheItem[T]) error) error {
	return rc.currentCache.ForEach(callback)
}

// Clean cleans the cache.
func (rc *ResizableCache[T]) Clean(expiringDate time.Time) int {
	rc.ResizeIf()
	return 0
}

// CleanParallel cleans the cache in parallel.
func (rc *ResizableCache[T]) CleanParallel(expiringDate time.Time) int {
	rc.ResizeIf()
	return 0
}

// ResizeIf resizes the cache if needed
func (rc *ResizableCache[T]) ResizeIf() bool {
	if rc.oldCache != nil {
		return false
	}

	//If false set to true
	if rc.isResizing.CAS(true, true) {
		return false
	}
	defer rc.isResizing.Store(false)

	count, capacity := rc.currentCache.CountCapacity()
	newSize, resize := rc.settings.ShouldResize(count, capacity)

	if !resize {
		return false
	}

	rc.settings.Capacity = int(newSize)
	newCache := NewFixedCache[T](rc.settings.FixedCacheSettings)

	//Make sure the oldCache is cleared
	old := rc.currentCache
	defer func() {
		rc.oldCache = nil
		old.Dispose()
	}()
	//Copy the current cache to the old cache
	rc.oldCache = rc.currentCache
	//Set the new cache as the current cache
	rc.currentCache = newCache

	now := time.Now()
	old.ForEach(func(value CacheItem[T]) error {
		if !value.IsExpired(now) {
			newCache.Set(value.GetValue())
		}
		return nil
	})

	return true
}
