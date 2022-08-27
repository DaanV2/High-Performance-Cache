package reports

import (
	"errors"
	"io/ioutil"
	"strings"
)

func (report *MarkDownReport) ParseTextFile(filepath string) error {
	data, err := ioutil.ReadFile(filepath,)
	if err != nil {
		return err
	}

	txt := string(data)
	lines := strings.Split(txt, "\n")

	benchmarkData := report.NewBenchmark()
	for _, line := range lines {
		//Result line
		if strings.Contains(line, ": ") {
			if err := benchmarkData.ParseAttribute(line); err != nil {
				return err
			}

		} else if strings.HasPrefix(line, "Benchmark") {
			if err := benchmarkData.ParseResult(line); err != nil {
				return err
			}

		} else if strings.HasPrefix(line, "Pass") || strings.HasPrefix(line, "Fail") {
			//DONE with this benchmark
			benchmarkData = report.NewBenchmark()
		}
	}

	return nil
}

const WhiteSpace =  " \t\n\r"

func (data *BenchMarkData) ParseAttribute(text string) error {
	index := strings.Index(text, ":")

	if index < 0 {
		return errors.New("Invalid attribute: " + text)
	}

	key := strings.Trim(text[:index], WhiteSpace)
	value := strings.Trim(text[index+1:], WhiteSpace)

	data.Attributes[key] = value

	return nil
}

func (data *BenchMarkData) ParseResult(text string) error {
	parts := strings.Split(text, "\t")

	if len(parts) < 2 {
		return errors.New("Invalid result: " + text)
	}

	result := data.NewResult()

	Name := parts[0][9:]
	Name = strings.ReplaceAll(Name, "_", " ")
	Name = strings.ReplaceAll(Name, "/", " ")
	result.Name = strings.Trim(Name, WhiteSpace)

	for index, part := range parts[1:] {
		part = strings.Trim(part, WhiteSpace)

		splitIndex := strings.Index(part, " ")
		if splitIndex >= 0 {
			unit := strings.Trim(part[splitIndex+1:], WhiteSpace)
			part = strings.Trim(part[:splitIndex], WhiteSpace)
			data.SetHeader(unit, index)
		} else {
			data.SetHeader("", index)
		}

		result.AddValue(part)
	}

	return nil
}