package concurrency

import "sync"

type Locks struct {
	locks []*sync.Mutex
}

func NewLocks(size int) Locks {
	result := Locks{
		locks: make([]*sync.Mutex, size),
	}

	for i := 0; i < size; i++ {
		result.locks[i] = &sync.Mutex{}
	}

	return result
}

func (locks Locks) GetLock(hashcode uint64) *sync.Mutex {
	index := int(hashcode) % len(locks.locks)

	return locks.locks[index]
}