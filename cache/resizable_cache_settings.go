package cache

import "time"

//ResizableCacheSettings is the settings for a resizable cache
type ResizableCacheSettings struct {
	FixedCacheSettings
	//ShouldResize is a function that returns true if the cache should be resized, and a value to how much
	ShouldResize func(count, capacity uint64) (uint64, bool)
	// ResizeClean is the cache cleaning settings of the resize cache alone
	ResizeClean CacheCleaningSettings
}

// DefaultResizableCacheSettings is the default settings for a resizable cache
func DefaultResizableCacheSettings[T KeyedObject](capacity int) ResizableCacheSettings {
	resizeSettings := DefaultCacheCleaningSettings()
	resizeSettings.Interval = 1 * time.Minute

	return ResizableCacheSettings{
		ShouldResize:       DefaultResizeFunction,
		FixedCacheSettings: DefaultFixedCacheSettings[T](capacity),
		ResizeClean:        resizeSettings,
	}
}

// DefaultResizeFunction is the default function that determines if the cache should be resized
func DefaultResizeFunction(count, capacity uint64) (uint64, bool) {
	// If the cache is 25% or lower empty, resize it to 25%
	lowerLimit := capacity / 4
	// If the cache is 75% or higher full, resize it to 200%
	upperLimit := (capacity * 3) / 4

	if count > upperLimit {
		return capacity * 2, true
	}

	if count < lowerLimit {
		return lowerLimit, true
	}

	return 0, false
}
