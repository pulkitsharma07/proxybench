package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

// testResultField has the actual result for example:
// "duration": 2.33
type testResultField struct {
	Name  string
	Value float64 // TODO: make this generic
}

func (tr testResultField) String() string {
	return fmt.Sprintf("(%s : %f)", tr.Name, tr.Value)
}

// testResult has many testResultFields
type testResult interface {
	testCase() testCase
	Results() []testResultField

	// Should pretty print the testName and results
	String() string
}

//LoadTimeResult is a testResult
type LoadTimeResult struct {
	resultOfTest testCase
	results      []testResultField
}

func NewLoadTimeResult(test testCase, name string, duration float64) LoadTimeResult {
	return LoadTimeResult{
		test,
		[]testResultField{{
			name,
			duration,
		}},
	}
}

func (l LoadTimeResult) testCase() testCase {
	return l.resultOfTest
}

func (l LoadTimeResult) Results() []testResultField {
	return l.results
}

func (l LoadTimeResult) String() string {
	return fmt.Sprintf("Result for testCase: %v, %v", l.testCase(), l.Results())
}

//TODO:
// Make this accept testResult instead.
func prettyPrint(results []LoadTimeResult) {
	data := [][]string{}

	for _, result := range results {
		resultFields := result.Results()
		data = append(data, []string{result.testCase().String(), resultFields[0].Name, fmt.Sprintf("%f", resultFields[0].Value)})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Benchmark", "Field", "Value"})

	table.SetHeaderColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.BgGreenColor},
	)

	table.SetColumnColor(
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiBlackColor},
		tablewriter.Colors{tablewriter.Bold, tablewriter.FgHiRedColor},
	)

	table.AppendBulk(data)
	fmt.Print("\nRESULTS\n")
	table.Render()
}
