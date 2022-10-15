package data

import (
	"errors"
	"strconv"
	"strings"
)

type TestCollector struct {
	Data map[string]*TestData `json:"data"`
}

func (collector *TestCollector) NewCollector() *TestCollector {
	return &TestCollector{}
}

func (collector *TestCollector) ParseTest(log string) error {
	//Format: <TestName>/<parameter>:<value>[,<parameter>:<value>]/<noise>\t[result]

	result := &TestResult{}

	index := strings.Index(log, "\t")
	if index < 0 {
		return errors.New("Cannot find name")
	}
	result.Name = log[:index]
	log = log[index+1:]

	index = strings.Index(log, "/")
	if index < 0 {
		return errors.New("Cannot find type")
	}
	result.Type = log[:index]
	log = log[index+1:]

	data := strings.Split(log, "\t")
	data = data[1:]

	for _, item := range data {
		item = strings.Trim(item, " \t")

		index = strings.Index(item, " ")
		if index < 0 {
			index = len(item)
		}

		valueStr := item[index:]
		unit := item[:index]
		value, err := strconv.ParseInt(valueStr, 10, 64)
		if err != nil {
			return err
		}

		switch (unit) {
		case "":
			result.Runs = value
		case "ns/op":
			result.Nanos = value
		case "items":
			result.TestSize = value
		case "B/op":
			result.Bytes = value
		case "allocs/op":
			result.Allocs = value
		}
	}

	if result.Valid() {
		result.Sanitize()
		collector.Add(result)
	}

	return nil
}

func (collector *TestCollector) Add(result *TestResult) {
	data, ok := collector.Data[result.Name]
	if !ok {
		data = &TestData{Name: result.Name}
		collector.Data[result.Name] = data
	}
	data.Add(result)
}
