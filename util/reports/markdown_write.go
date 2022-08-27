package reports

import (
	"fmt"
	"os"
	"sort"
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

		//Write test attributes
		if len(report.Attributes) > 0 {
			writer.WriteString("|Test Attributes|Value|\n")
			writer.WriteString("|---------------|:-----|\n")
			attributes := make([]string, len(report.Attributes))

			for key, value := range report.Attributes {
				attributes = append(attributes, fmt.Sprintf("|%s|%s|", key, value))
			}

			sort.SliceStable(attributes, func(i, j int) bool {
				return strings.Compare(attributes[i], attributes[j]) < 0
			})

			for _, attribute := range attributes {
				writer.WriteString(attribute)
			}
			writer.WriteString("\n")
		}

		//Write title
		if _, err := writer.WriteString(fmt.Sprintf("## Benchmark: %s \n\n", report.Title)); err != nil {
			return err
		}

		//Write test results
		writer.WriteString("|Name|" + strings.Join(report.Headers, "|") + "|\n")
		writer.WriteString("|----:|" + strings.Repeat("---:|", len(report.Headers)) + "\n")

		for _, result := range report.Results {
			writer.WriteString("|" + result.Name + "|" + strings.Join(result.Values, "|") + "|\n")
		}
		writer.WriteString("\n")
	}

	return nil
}
