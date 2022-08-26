package cache

import "math"

type HashRange struct {
	minHash int32
	maxHash int32
}

func NewHashRange() HashRange {
	return HashRange{
		minHash: math.MaxInt32,
		maxHash: math.MinInt32,
	}
}

func (hr HashRange) IsInRange(hashcode int32) bool {
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
func (hr *HashRange) UpdateRange(hashcode int32) {
	if hashcode < hr.minHash {
		hr.minHash = hashcode
	}
	if hashcode > hr.maxHash {
		hr.maxHash = hashcode
	}
}
