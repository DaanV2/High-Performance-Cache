package benchmarks

import (
	"fmt"
	"testing"
)

func Benchmark_Duplicate_Items_Write_Test(b *testing.B) {
	temp := DataSet
	data := make([]*BenchmarkData, len(temp))

	for i := 0; i < len(temp); i++ {
		data[i] = temp[i % 64]
	}

	tests := []*TestSettings{
		{
			Name: "MapCache",
			CreateCache: newMapCache[*BenchmarkData],
		},
	}

	RunBenchMarks(b, data, tests, WriteTest)

	fmt.Println("Done")
}

func Benchmark_Duplicate_Items_Read_Test(b *testing.B) {
	data := DataSet

	tests := []*TestSettings{
		{
			Name: "MapCache",
			CreateCache: newMapCache[*BenchmarkData],
		},
	}

	RunBenchMarks(b, data, tests, ReadTest)

	fmt.Println("Done")
}



