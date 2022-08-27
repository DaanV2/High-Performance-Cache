package benchmarks

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"time"
)

const TestDataSize = 100_000
const PowerStep = 10

func init() {
	DataSet = make([]*BenchmarkData, TestDataSize)

	fmt.Printf("Generating data set with %v items\n", len(DataSet))
	max := uint64(len(DataSet))

	for i := uint64(0); i < max; i++ {
		DataSet[i] = NewBenchMarData(i)
	}
}

func GenerateSizes[T any](items []T) []int {
	Sizes := make([]int, 0)

	for size := PowerStep; size < len(DataSet); {
		Sizes = append(Sizes, size)
		size *= PowerStep
	}

	return Sizes
}

var DataSet []*BenchmarkData

type BenchmarkData struct {
	Id          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time
}

func NewBenchMarData(id uint64) *BenchmarkData {
	bytes := make([]byte, 8)
	binary.LittleEndian.PutUint64(bytes, id)
	str := hex.EncodeToString(bytes)
	str = str + str

	return &BenchmarkData{
		Id:          str,
		Name:        str,
		Description: str,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}
}

func (b *BenchmarkData) GetKey() string {
	return b.Id
}
