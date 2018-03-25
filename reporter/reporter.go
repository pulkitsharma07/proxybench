package reporter

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/pulkitsharma07/proxybench/result"
)

func ShowReport(results []result.Result) {
	data := [][]string{}

	for _, result := range results {
		data = append(data, []string{result.Config().String(), result.Benchmark(), fmt.Sprintf("%+v", result.Results())})
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Config", "Benchmark", "Result(s)"})

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
	fmt.Print("\nREPORT\n")
	table.Render()
}
