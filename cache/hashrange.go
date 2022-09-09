package cache

import (
	"fmt"
	"math"
)

type HashRange struct {
	minHash uint64
	maxHash uint64
}

// NewHashRange creates a new HashRange
func NewHashRange() HashRange {
	return NewHashRangeFrom(math.MaxUint64, 0)
}

// NewHashRangeFrom creates a new HashRange from the given min and max hashcode.
func NewHashRangeFrom(min, max uint64) HashRange {
	return HashRange{
		minHash: min,
		maxHash: max,
	}
}

// Contains returns true if the given hashcode is in the range.
func (hr HashRange) Contains(hashcode uint64) bool {
	if hashcode < hr.minHash {
		return false
	}
	if hashcode > hr.maxHash {
		return false
	}

	return true
}

// UpdateRange updates the range with the given hashcode.
// If the hashcode is outside the range, the range is expanded.
// This has to be done by reference, because the range is updated.
func (hr *HashRange) UpdateRange(hashcode uint64) {
	if hashcode < hr.minHash {
		hr.minHash = hashcode
	}
	if hashcode > hr.maxHash {
		hr.maxHash = hashcode
	}
}

// Update updates the range with the given range.
func (hr *HashRange) Update(r HashRange) {
	hr.UpdateRange(r.minHash)
	hr.UpdateRange(r.maxHash)
}

//Clear clears the range.
func (hr *HashRange) Clear() {
	hr.minHash = math.MaxUint64
	hr.maxHash = 0
}

func (hr *HashRange) String() string {
	return fmt.Sprintf("[%d, %d]", hr.minHash, hr.maxHash)
}
