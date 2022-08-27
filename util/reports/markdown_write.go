package reports

import (
	"fmt"
	"os"
	"strings"
)

func (report *MarkDownReport) WriteTo(filepath string) error {
	writer, err := os.Create(filepath)
	defer writer.Close()

	if err != nil {
		return err
	}

	writer.WriteString("# Benchmark results\n\n")

	for _, report := range report.Benchmarks {
		if !report.HasData() {
			continue
		}

		//Write title
		if _, err := writer.WriteString(fmt.Sprintf("## Benchmark: %s \n\n", report.Title)); err != nil {
			return err
		}

		//Write test attributes
		if len(report.Attributes) > 0 {
			writer.WriteString("|Test Attributes|Value|\n")
			writer.WriteString("|---------------|-----|\n")
			for key, value := range report.Attributes {
				writer.WriteString(fmt.Sprintf("|%v|%v|\n", key, value))
			}
			writer.WriteString("\n")
		}

		//Write test results
		writer.WriteString("|Name|" + strings.Join(report.Headers, "|") + "|\n")
		writer.WriteString("|----|" + strings.Repeat("---|", len(report.Headers)) + "\n")

		for _, result := range report.Results {
			writer.WriteString("|" + result.Name + "|" + strings.Join(result.Values, "|") + "|\n")
		}
		writer.WriteString("\n")
	}

	return nil
}
