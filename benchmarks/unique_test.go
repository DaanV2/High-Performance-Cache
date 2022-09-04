package benchmarks

import (
	"fmt"
	"testing"
)

func Benchmark_Unique_Items_Write_Test(b *testing.B) {
	data := DataSet
	tests := CachesToTest

	RunBenchMarks(b, data, tests, WriteTest)

	fmt.Println("Done")
}

func Benchmark_Unique_Items_Read_Test(b *testing.B) {
	data := DataSet
	tests := CachesToTest

	RunBenchMarks(b, data, tests, ReadTest)

	fmt.Println("Done")
}
