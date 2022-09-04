package cache

import "github.com/DaanV2/High-Performance-Cache/util"

type KeyLookup struct {
	Hashcode uint64
	Key      string
}

func NewKeyLookup(key string) KeyLookup {
	return KeyLookup{
		Hashcode: util.GetHashcode(key),
		Key:      key,
	}
}

func (c KeyLookup) GetKey() string {
	return c.Key
}

func (c KeyLookup) GetHashcode() uint64 {
	return c.Hashcode
}