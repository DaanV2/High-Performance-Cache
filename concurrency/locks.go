package concurrency

import "sync"

// Locks is a set of locks.
type Locks struct {
	// the set of locks.
	locks []*sync.Mutex
}

// NewLocks returns a new set of locks.
func NewLocks(size int) Locks {
	result := Locks{
		locks: make([]*sync.Mutex, size),
	}

	for i := 0; i < size; i++ {
		result.locks[i] = &sync.Mutex{}
	}

	return result
}

// GetLock returns the lock for the given hashcode.
func (locks Locks) GetLock(hashcode uint64) *sync.Mutex {
	index := int(hashcode) % len(locks.locks)

	return locks.locks[index]
}

func (locks Locks) GetLockByIndex(index int) *sync.Mutex {
	return locks.locks[index]
}

// Count returns the number of locks.
func (locks Locks) Count() int {
	return len(locks.locks)
}

// Lock locks the lock for the given hashcode.
func (locks Locks) Lock(hashcode uint64) {
	locks.GetLock(hashcode).Lock()
}

// TryLock tries to lock the lock for the given hashcode.
func (locks Locks) TryLock(hashcode uint64) bool {
	return locks.GetLock(hashcode).TryLock()
}

// Unlock unlocks the lock for the given hashcode.
func (locks Locks) Unlock(hashcode uint64) {
	locks.GetLock(hashcode).Unlock()
}

// LockAll locks all locks.
func (locks Locks) LockAll() {
	for _, lock := range locks.locks {
		lock.Lock()
	}
}

// UnlockAll unlocks all locks.
func (locks Locks) UnlockAll() {
	for _, lock := range locks.locks {
		lock.Unlock()
	}
}