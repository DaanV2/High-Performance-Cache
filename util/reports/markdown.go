package reports

import "errors"

type MarkDownReport struct {
	BenchMarks []*BenchMarkData
}

type BenchMarkData struct {
	Attributes map[string]string
	Results    []*BenchMarkResults
	Headers    []string
}

type BenchMarkResults struct {
	Name   string
	Values []string
}

func NewMarkDownReport() *MarkDownReport {
	return &MarkDownReport{
		BenchMarks: []*BenchMarkData{},
	}
}

func (report *MarkDownReport) NewBenchmark() *BenchMarkData {
	data := &BenchMarkData{
		Attributes: map[string]string{},
		Results:    []*BenchMarkResults{},
		Headers:    []string{},
	}
	report.BenchMarks = append(report.BenchMarks, data)

	return data
}

func (report *BenchMarkData) NewResult() *BenchMarkResults {
	result := &BenchMarkResults{
		Values: []string{},
	}
	report.Results = append(report.Results, result)

	return result
}

func (report *BenchMarkData) HasData() bool {
	return len(report.Results) > 0
}

func (report *BenchMarkData) SetHeader(unit string, index int) error {
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

func (data *BenchMarkResults) AddValue(value string) {
	data.Values = append(data.Values, value)
}