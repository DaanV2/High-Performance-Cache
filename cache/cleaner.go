package cache

import (
	"context"
	"time"

	"go.uber.org/zap"
)

type CacheCleaningSettings struct {
	//AutoClean determines if the cache should automatically clean itself
	AutoClean bool
	//CleanInterval is the interval at which the cache should clean itself
	Interval time.Duration
	//Parralel determines if the cleaning should be done in parralel
	Parralel bool
	//Logger is the logger that should be used for logging
	Logger *zap.Logger
}

func DefaultCacheCleaningSettings() CacheCleaningSettings {
	return CacheCleaningSettings{
		AutoClean:  true,
		Interval:   time.Minute * 10,
		Parralel:   true,
		Logger: zap.L(),
	}
}

type CacheCleaner struct {
	settings CacheCleaningSettings
	cache    Cleanable
	ctx      context.Context
	close    context.CancelFunc
}

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

func (cc *CacheCleaner) Dispose() {
	cc.close()
	cc.cache = nil
}

func (cc *CacheCleaner) Start() {
	go func() {
		cc.settings.Logger.Info("Starting cache cleaner")
		cc.loop()
		cc.settings.Logger.Info("Stopping cache cleaner")
	}()
}

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
		if cc.settings.Parralel {
			cc.settings.Logger.Info("Cleaning cache in parralel...")
			amount = cache.CleanParralel(now)
		} else {
			cc.settings.Logger.Info("Cleaning cache...")
			amount = cache.Clean(now)
		}

		cc.settings.Logger.Info("Cleaned cache", zap.Int("amount", amount))
	}
}
