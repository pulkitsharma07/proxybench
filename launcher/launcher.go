package launcher

import (
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/suite"
)

func Launch(suiteConfig []config.Config) *suite.Suite {
	Suite := suite.NewSimpleSuite(suiteConfig)
	Suite.Run()

	return Suite
}
