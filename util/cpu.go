package util

import (
	util_os "github.com/DaanV2/High-Performance-Cache/util/os"
	"github.com/klauspost/cpuid/v2"
)

func init() {
	CPU := cpuid.CPU

	cpuinfo = &CPUData{
		BrandName: CPU.BrandName,
		Cache: &CacheInfo{
			L1: util_os.GetEnvironmentVariable("CPU_CACHE_L1", util_os.ToInt64, int64(CPU.Cache.L1I)),
			L2: util_os.GetEnvironmentVariable("CPU_CACHE_L2", util_os.ToInt64, int64(CPU.Cache.L1I)),
			L3: util_os.GetEnvironmentVariable("CPU_CACHE_L3", util_os.ToInt64, int64(CPU.Cache.L1I)),
		},
	}

	cpuinfo.Cache.CheckAndEstimate()

}

var cpuinfo *CPUData

func GetCPUInfo() *CPUData {
	return cpuinfo
}

type CPUData struct {
	BrandName string
	Cache     *CacheInfo
}

type CacheInfo struct {
	L1 int64
	L2 int64
	L3 int64
}

func (c *CacheInfo) CheckAndEstimate() {

	if c.L1 <= 0 {
		c.L1 = 128_000 // 128KB is currently the default value for L1 cache size of a modern CPU
	}
	if c.L2 <= 0 {
		c.L2 = 8_388_608 // 8MB is currently the default value for L2 cache size of a modern CPU
	}
	if c.L3 <= 0 {
		c.L3 = 33_554_432 // 32MB is currently the default value for L3 cache size of a modern CPU
	}
}
