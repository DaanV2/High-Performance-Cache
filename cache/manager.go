package cache

type CacheManagerSettings struct {
	Cleaning CacheCleaningSettings
}

type CacheManager struct {
	settings CacheManagerSettings
	cleaner  *CacheCleaner
}

func NewManager(settings CacheManagerSettings) *CacheManager {
	return &CacheManager{
		settings: settings,
		cleaner:  StartCacheCleaner(settings.Cleaning),
	}
}

// Add adds a cache to the manager, will also be registered with the cleaner
func (manager *CacheManager) Add(cache Cleanable) {
	manager.cleaner.Push(cache)
}

// Remove removes a cache from the manager, will also be unregistered with the cleaner
func (manager *CacheManager) Remove(cache Cleanable) {
	manager.cleaner.Remove(cache)
}

// Dispose disposes the manager, but none of the caches
func (manager *CacheManager) Dispose() {
	manager.cleaner.Dispose()
}

// DisposeAll disposes the manager and all of the caches
func (manager *CacheManager) DisposeAll() {
	manager.Dispose()
	caches := manager.cleaner.caches

	for _, cache := range caches {
		if t, ok := cache.(Disposable); ok {
			t.Dispose()
		}
	}
}

// ForEach iterates over all caches
func (manager *CacheManager) ForEach(f func(cache Cleanable)) {
	caches := manager.cleaner.caches

	for _, cache := range caches {
		f(cache)
	}
}
