package main

func main() {
	chromeDriver := Driver{"chromedriver", 9222, nil}
	chromeDriver.Start()
	defer chromeDriver.Stop()

	suite := NewSimpleTestSuite(chromeDriver)
	suite.executeAsync()

	//executeSync(tests, chromeDriver)
}
