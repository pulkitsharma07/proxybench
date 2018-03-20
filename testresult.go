package main

// testResultField has the actual result for example:
// "duration": 2.33
type testResultField struct {
	Name  string
	Value float64 // TODO: make this generic
}

// testResult has many testResultFields
type testResult interface {
	testCase() *testCase
	Results() []testResultField
}

//LoadTimeResult is a testResult
type LoadTimeResult struct {
	resultOfTest *testCase
	results      []testResultField
}

func NewLoadTimeResult(test testCase, name string, duration float64) *LoadTimeResult {
	return &LoadTimeResult{
		&test,
		[]testResultField{{
			name,
			duration,
		}},
	}
}

func (l *LoadTimeResult) testCase() *testCase {
	return l.resultOfTest
}

func (l *LoadTimeResult) Results() []testResultField {
	return l.results
}
