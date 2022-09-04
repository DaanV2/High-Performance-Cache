package benchmarks

import (
	"fmt"
	"testing"
)

func Benchmark_Duplicate_Items_Write_Test(b *testing.B) {
	temp := DataSet
	data := make([]*BenchmarkData, len(temp))
	tests := CachesToTest

	for i := 0; i < len(temp); i++ {
		data[i] = temp[i%64]
	}

	RunBenchMarks(b, data, tests, WriteTest)

	fmt.Println("Done")
}

func Benchmark_Duplicate_Items_Read_Test(b *testing.B) {
	temp := DataSet
	data := make([]*BenchmarkData, len(temp))
	tests := CachesToTest

	for i := 0; i < len(temp); i++ {
		data[i] = temp[i%64]
	}

	RunBenchMarks(b, data, tests, ReadTest)

	fmt.Println("Done")
}
