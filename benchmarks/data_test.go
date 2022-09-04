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

func Sizes() []int {
	return []int{
		100,
		200,
		500,
		1000,
		2000,
		5000,
		10000,
		20000,
		50000,
		100000,
	}
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
