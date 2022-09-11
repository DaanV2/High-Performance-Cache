package cache

import "github.com/DaanV2/High-Performance-Cache/util"

// KeyLookup is a lookup for a key
type KeyLookup struct {
	//Hashcode is the hashcode of the key
	Hashcode uint64
	//Key is the key
	Key string
}

// NewKeyLookup creates a new KeyLookup
func NewKeyLookup(key string) KeyLookup {
	return KeyLookup{
		Hashcode: util.GetHashcode(key),
		Key:      key,
	}
}

// GetKey returns the key
func (c KeyLookup) GetKey() string {
	return c.Key
}

// GetHashcode returns the hashcode of the key
func (c KeyLookup) GetHashcode() uint64 {
	return c.Hashcode
}
