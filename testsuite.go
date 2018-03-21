package main

import (
	"fmt"
	"time"

	"github.com/pulkitsharma07/proxybench/benchmark"
)

type testSuite struct {
	config testConfig
	driver Driver
	tests  []benchmark.Benchmark
}

// Generates a new testSuite, Add different tests in the testCase slice
func NewSimpleTestSuite(driver Driver) *testSuite {
	return &testSuite{
		testConfig{
			Proxy{},
		},
		driver,
		[]benchmark.Benchmark{
			StressHTTP("HTTP Stress"),
			StressHTTPS("HTTPS Stress"),
		},
	}
}

func (t *testSuite) String() string {
	return fmt.Sprintf("ProxyConfig: %v, Driver: %v, tests: %v \n", t.config, t.driver, t.tests)
}

func (t *testSuite) executeAsync() {
	fmt.Printf("Executing Test Suite(Async)\nPlease Wait..\n")

	// Slice to capture the results for printing report later.
	results := []LoadTimeResult{}

	// channel for gathering results
	ch := make(chan testResult)

	testsToExecute := t.tests

	// Launch tests parallely
	for _, test := range testsToExecute {
		//fmt.Printf("****Starting %v*****\n", test)
		go test.Run(t.config, t.driver.port, ch)

		// Do not fire concurrent start requests to chromedriver.
		time.Sleep(1 * time.Second)
	}

	for range testsToExecute {
		var res = <-ch
		results = append(results, res.(LoadTimeResult))
		fmt.Printf("Completed(%d/%d)\n", len(results), len(testsToExecute))
	}

	prettyPrint(results)
}

func (t *testSuite) executeSync() {
	fmt.Printf("Executing Test Suite(Sync)\nPlease Wait..\n")

	// Slice to capture the results for printing report later.
	results := []LoadTimeResult{}

	// channel for gathering results
	ch := make(chan testResult)

	testsToExecute := t.tests

	// Launch tests in Sync
	for _, test := range testsToExecute {
		//fmt.Printf("****Starting %v*****\n", test)
		go test.Run(t.config, t.driver.port, ch)

		var res = <-ch
		results = append(results, res.(LoadTimeResult))
		fmt.Printf("Completed(%d/%d)\n", len(results), len(testsToExecute))
	}
	prettyPrint(results)
}
