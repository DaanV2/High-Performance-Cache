package reports

import (
	"errors"
	"strings"
)

type MarkDownReport struct {
	Benchmarks []*BenchmarkData
}

type BenchmarkData struct {
	Title      string
	Attributes map[string]string
	Results    []*BenchmarkResults
	Headers    []string
}

type BenchmarkResults struct {
	Name   string
	Values []string
}

func NewMarkDownReport() *MarkDownReport {
	return &MarkDownReport{
		Benchmarks: []*BenchmarkData{},
	}
}

func (report *MarkDownReport) NewBenchmark() *BenchmarkData {
	data := &BenchmarkData{
		Attributes: map[string]string{},
		Results:    []*BenchmarkResults{},
		Headers:    []string{},
	}
	report.Benchmarks = append(report.Benchmarks, data)

	return data
}

func (report *BenchmarkData) NewResult() *BenchmarkResults {
	result := &BenchmarkResults{
		Values: []string{},
	}
	report.Results = append(report.Results, result)

	return result
}

func (report *BenchmarkData) HasData() bool {
	return len(report.Results) > 0
}

func (report *BenchmarkData) SetHeader(unit string, index int) error {
	for index >= len(report.Headers) {
		report.Headers = append(report.Headers, "")
	}

	header := report.Headers[index]
	if header != "" {
		return nil
	}

	switch unit {
	case "ns/op":
		header = "Nanoseconds per operation"
	case "B/op":
		header = "Bytes per operation"
	case "allocs/op":
		header = "Allocations per operation"
	case "":
		header = "Test Amount"
	default:
		return errors.New("Invalid unit: " + unit)
	}

	report.Headers[index] = header
	return nil
}

func (report *BenchmarkData) DetermineTitle() {
	if len(report.Results) == 0 {
		return
	}

	var count int
	for count = 0; report.sameCharacterInName(count); count++ {
	}

	count = strings.LastIndex(report.Results[0].Name[:count], " ")
	title := report.Results[0].Name[:count]
	report.Title = strings.Trim(title, WhiteSpace)

	for index, result := range report.Results {
		result.Name = strings.Trim(result.Name[count:], WhiteSpace)
		report.Results[index] = result
	}
}

func (report *BenchmarkData) sameCharacterInName(index int) bool {
	char := report.Results[0].Name[index]

	for _, result := range report.Results[1:] {
		if result.Name[index] != char {
			return false
		}
	}

	return true
}

func (data *BenchmarkResults) AddValue(value string) {
	data.Values = append(data.Values, value)
}
