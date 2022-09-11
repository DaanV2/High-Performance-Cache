package cache

import (
	"time"

	"go.uber.org/zap"
)

// CacheCleaningSettings is the settings for a cache cleaner
type CacheCleaningSettings struct {
	//AutoClean determines if the cache should automatically clean itself
	AutoClean bool
	//CleanInterval is the interval at which the cache should clean itself
	Interval time.Duration
	//Parallel determines if the cleaning should be done in parallel
	Parallel bool
	//Logger is the logger that should be used for logging
	Logger *zap.Logger
}

// DefaultCacheCleaningSettings is the default settings for a cache cleaner
func DefaultCacheCleaningSettings() CacheCleaningSettings {
	return CacheCleaningSettings{
		AutoClean: true,
		Interval:  time.Minute * 10,
		Parallel:  true,
		Logger:    zap.L().With(zap.String("type", "cacheCleaner")),
	}
}
