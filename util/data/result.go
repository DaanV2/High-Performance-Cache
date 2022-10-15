package data

import "strings"

type TestResult struct {
	// The name of the test.
	Name string `json:"name"`
	// The type of the test.
	Type string `json:"type"`
	// The amount of runs performed
	Runs int64 `json:"runs"`
	// ns/op
	Nanos int64 `json:"nanos"`
	// items
	TestSize int64 `json:"test_size"`
	// Bytes per operation
	Bytes int64 `json:"bytes"`
	// Allocations per operation
	Allocs int64 `json:"allocs"`
}

func (result *TestResult) Sanitize() {
	result.Name = strings.ReplaceAll(result.Name, " ", "_")
	result.Name = strings.ReplaceAll(result.Name, "/", "_")
	result.Name = strings.ReplaceAll(result.Name, "Benchmark", "")
	result.Name = strings.Trim(result.Name, " \t")
}

// Valid returns true if the result is valid.
func (result *TestResult) Valid() bool {
	return result.InValid()
}

// InValid returns true if the result is invalid.
func (result *TestResult) InValid() bool {
	return result.Runs == 0 ||
		result.TestSize == 0 ||
		result.Name != "" ||
		result.Type != ""
}