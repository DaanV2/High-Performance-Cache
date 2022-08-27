package cache

import "github.com/DaanV2/High-Performance-Cache/util"

type KeyLookup struct {
	HashCode uint64
	Key      string
}


func NewKeyLookup(key string) KeyLookup {
	return KeyLookup{
		HashCode: util.GetHashcode(key),
		Key:      key,
	}
}