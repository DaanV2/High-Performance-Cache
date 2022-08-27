package benchmarks

import (
	"fmt"
	"testing"
)

func Benchmark_Default_Write_Test(b *testing.B) {
	data := DataSet

	tests := []*TestSettings{
		{
			Name: "MapCache",
			CreateCache: newMapCache[*BenchmarkData],
		},
	}

	RunBenchMarks(b, data, tests, WriteTest)

	fmt.Println("Done")
}

func Benchmark_Default_Read_Test(b *testing.B) {
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



func Name(name string, a ...any) string {
	return fmt.Sprintf(name, a...)
}
