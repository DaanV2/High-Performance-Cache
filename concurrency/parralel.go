package concurrency

import (
	"runtime"
	"sync"
)

var MaxConcurrency = runtime.GOMAXPROCS(0)

func Parralel[T any](items []T, call func(item T, index int, current []T)) {
	max := MaxConcurrency
	group := &sync.WaitGroup{}
	count := len(items)
	step := count / max

	for i := 0; i < count; i += step {
		group.Add(1)
		max := i+step
		if max > count {
			max = count
		}
		go parralelSet(items[i:max], group, call)
	}

	group.Wait()
}

func parralelSet[T any](items []T, group *sync.WaitGroup, call func(item T, index int, current []T)) {
	for i, item := range items {
		call(item, i, items)
	}

	group.Done()
}
