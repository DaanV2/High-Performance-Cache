package data

type TestData struct {
	Name   string                 `json:"name"`
	Series map[string]*TestSeries `json:"series"`
}

func (data *TestData) Add(result *TestResult) {
	series, ok := data.Series[result.Type]
	if !ok {
		series = &TestSeries{
			Type: result.Type,
			Data: make([]*TestResult, 0),
		}
		data.Series[result.Type] = series
	}
	series.Add(result)
}
