package main

import (
	"fmt"
	"github.com/pulkitsharma07/proxybench/automation"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/suite"
)

func main() {
	chromeDriver := automation.NewChromeDriver(9222)
	chromeDriver.Start()
	defer chromeDriver.Stop()

	Suite := suite.NewSimpleSuite([]config.Config{{config.Proxy{}}}, chromeDriver.Port)

	fmt.Printf("%+v", Suite)
	Suite.Run()

	Suite.Results()
}
