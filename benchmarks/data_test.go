package benchmarks

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/DaanV2/High-Performance-Cache/util"
)

const TestDataSize = 100_000
const PowerStep = 10

func init() {
	DataSet = make([]*BenchmarkData, TestDataSize)

	fmt.Printf("Generating data set with %v items\n", len(DataSet))
	for i := 0; i < len(DataSet); i++ {
		DataSet[i] = NewBenchMarData()
	}

	fmt.Println("Generating data done")
	fmt.Printf("concurrency: %v\n", util.MaxConcurrency)
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

func NewBenchMarData() *BenchmarkData {
	return &BenchmarkData{
		Id:          RandomStr(16),
		Name:        RandomStr(16),
		Description: RandomStr(16),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		DeletedAt:   time.Now(),
	}
}

func (b *BenchmarkData) GetKey() string {
	return b.Id
}

func RandomStr(length int) string {
	bytes := make([]byte, length)
	rand.Read(bytes)

	return hex.EncodeToString(bytes)
}
