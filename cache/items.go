package cache

type CacheItemString string

func (c *CacheItemString) GetKey() string {
	return string((*c))
}