package util

import "hash/fnv"

func GetHashcode(key string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(key))
	return h.Sum64()
}