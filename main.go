package main

import (
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/launcher"
	"github.com/pulkitsharma07/proxybench/reporter"
)

func main() {
	suite := launcher.Launch(config.ReadConfig())
	reporter.ShowReport(suite.Results())
}
