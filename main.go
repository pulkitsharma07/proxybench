package main

import (
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/launcher"
	"github.com/pulkitsharma07/proxybench/reporter"
)

func main() {
	suite := launcher.Launch([]config.Config{
		{config.Proxy{"Direct", ""}},
		{config.Proxy{"mitmproxy", "localhost:8085"}},
		{config.Proxy{"browsermob (Legacy)", "localhost:8081"}},
		{config.Proxy{"browsermob (LittleProxy)", "localhost:8083"}},
	})
	reporter.ShowReport(suite.Results())
}
