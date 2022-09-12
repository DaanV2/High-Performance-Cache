package cache

import (
	"context"
	"time"

	"go.uber.org/zap"
)

// CacheCleaner is a cache cleaner
type CacheCleaner struct {
	// settings is the settings of the cache cleaner
	settings CacheCleaningSettings
	// cache is the cache that should be cleaned
	cache Cleanable
	// ctx is the context of the cache cleaner
	ctx context.Context
	// close is the function that should be called to close the cache cleaner
	close context.CancelFunc
}

// StartCacheCleaner starts a cache cleaner
func StartCacheCleaner(settings CacheCleaningSettings, cache Cleanable) *CacheCleaner {
	ctx, close := context.WithCancel(context.Background())

	result := &CacheCleaner{
		settings: settings,
		cache:    cache,
		ctx:      ctx,
		close:    close,
	}

	if result.settings.AutoClean {
		result.Start()
	}

	return result
}

// Stop stops the cache cleaner
func (cc *CacheCleaner) Dispose() {
	defer func() {
		cc.cache = nil
	}()
	cc.close()
}

// Start starts the cache cleaner
func (cc *CacheCleaner) Start() {
	go func() {
		cc.settings.Logger.Info("Starting cache cleaner")
		cc.loop()
		cc.settings.Logger.Info("Stopping cache cleaner")
	}()
}

// Stop stops the cache cleaner
func (cc *CacheCleaner) Stop() {
	cc.close()
}

// loop loops the cache cleaner
func (cc *CacheCleaner) loop() {
	for {
		select {
		case <-cc.ctx.Done():
			return
		case <-time.After(cc.settings.Interval):
		}

		cache := cc.cache
		if cache == nil {
			cc.settings.Logger.Warn("Cache is nil...")
			return
		}

		now := time.Now()
		amount := 0
		if cc.settings.Parallel {
			cc.settings.Logger.Info("Cleaning cache in parallel...")
			amount = cache.CleanParallel(now)
		} else {
			cc.settings.Logger.Info("Cleaning cache...")
			amount = cache.Clean(now)
		}

		cc.settings.Logger.Info("Cleaned cache", zap.Int("amount", amount))
	}
}
