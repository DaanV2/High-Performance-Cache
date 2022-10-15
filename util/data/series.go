package data

type TestSeries struct {
	Type string
	Data []*TestResult
}

func (series *TestSeries) Add(result *TestResult) {
	series.Data = append(series.Data, result)
}