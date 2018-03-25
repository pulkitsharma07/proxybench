package launcher

import (
	"github.com/pulkitsharma07/proxybench/automation"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/suite"
)

func Launch(suiteConfig []config.Config) *suite.Suite {
	chromeDriver := automation.NewChromeDriver(9222)
	chromeDriver.Start()
	defer chromeDriver.Stop()

	Suite := suite.NewSimpleSuite(suiteConfig, chromeDriver.Port)
	Suite.Run()

  return Suite
}
