package benchmarks

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"

	"github.com/DaanV2/High-Performance-Cache/util"
)

const TestDataSize = 100_000

func init() {
	DataSet = make([]*BenchmarkData, TestDataSize)
	Sizes = make([]int, 0)

	fmt.Printf("Generating data set with %v items\n", len(DataSet))
	for i := 0; i < len(DataSet); i++ {
		DataSet[i] = NewBenchMarData()
	}

	for size := 10; size < len(DataSet); {
		Sizes = append(Sizes, size)
		size *= 10
	}

	fmt.Println("Generating data done")
	fmt.Printf("concurrency: %v\n", util.MaxConcurrency)
}

var DataSet []*BenchmarkData
var Sizes []int

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
