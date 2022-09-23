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
	caches []Cleanable
	// ctx is the context of the cache cleaner
	ctx context.Context
	// close is the function that should be called to close the cache cleaner
	close context.CancelFunc
}

// StartCacheCleaner starts a cache cleaner
func StartCacheCleaner(settings CacheCleaningSettings) *CacheCleaner {
	ctx, close := context.WithCancel(context.Background())

	result := &CacheCleaner{
		settings: settings,
		caches:   make([]Cleanable, 0, 1),
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
		cc.caches = nil
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
			cc.Clean()
		}
	}
}

// Clean performs the cache cleaning
func (cc *CacheCleaner) Clean() {
	caches := cc.caches
	for _, cache := range caches {
		if cache == nil {
			cc.settings.Logger.Warn("Cache is nil...")
			continue
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

func (cc *CacheCleaner) Push(cache Cleanable) {
	cc.caches = append(cc.caches, cache)
}

func (cc *CacheCleaner) Remove(cache Cleanable) {
	for i, c := range cc.caches {
		if c == cache {
			cc.caches = append(cc.caches[:i], cc.caches[i+1:]...)
			return
		}
	}
}
