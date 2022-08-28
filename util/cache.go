package util

type CacheKind int

const (
	CacheL1 CacheKind = 0
	CacheL2 CacheKind = 1
	CacheL3 CacheKind = 2
)

// GetCacheSize returns the size of the cache.
func GetCacheSize(kind CacheKind) int64 {
	cacheInfo := GetCPUInfo().Cache

	switch kind {
	case CacheL1:
		return cacheInfo.L1
	case CacheL2:
		return cacheInfo.L2
	case CacheL3:
		return cacheInfo.L3
	default:
		return cacheInfo.L3
	}
}
