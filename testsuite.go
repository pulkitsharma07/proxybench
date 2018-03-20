package main

import (
	"fmt"
	"time"
)

type testSuite struct {
	config testConfig
	driver Driver
	tests  []testCase
}

// Generates a new testSuite, Add different tests in the testCase slice
func NewSimpleTestSuite(driver Driver) *testSuite {
	return &testSuite{
		testConfig{
			Proxy{},
		},
		driver,
		[]testCase{
			HTTPTestCase{
				"HTTP Stress",
			},
			HTTPSTestCase{
				"HTTPS Stress",
			},
		},
	}
}

func (t *testSuite) executeAsync() {
	fmt.Printf("Executing tests %+v parallely\n", t.tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests parallely
	for _, test := range t.tests {
		go test.Run(t.config, t.driver.port, ch)

		// Do not fire concurrent start requests to chromedriver.
		time.Sleep(1 * time.Second)
	}

	for range t.tests {
		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}

func (t *testSuite) executeSync() {
	fmt.Printf("Executing tests %+v sequentially\n", t.tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests in Sync
	for _, test := range t.tests {
		go test.Run(t.config, t.driver.port, ch)

		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}
