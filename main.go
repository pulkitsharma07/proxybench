package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

type testCase struct {
	url   string
	proxy Proxy
}

type testResult struct {
	test   testCase
	result float64 //TODO: Make this generic
}

func LaunchTest(test testCase, port int, done chan testResult) {
	b := &BrowserAutomator{port}

	wd := b.StartChrome(test.proxy)

	if err := wd.Get(test.url); err != nil {
		panic(err)
	}

	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	if err != nil {
		log.Fatal("Failed to parse results")
	}

	timeToLoadStr, err := timeToLoad.Text()
	duration, err := strconv.ParseFloat(strings.Replace(timeToLoadStr, " s", "", -1), 64)

	wd.Quit()
	done <- testResult{test, duration}
}

func executeAsync(tests []testCase, driver Driver) {
	fmt.Printf("Executing tests %+v parallely\n", tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests parallely
	for _, test := range tests {
		go LaunchTest(test, driver.portToListenAt, ch)

		// Do not fire concurrent start requests to chromedriver.
		time.Sleep(1 * time.Second)
	}

	for range tests {
		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}

func executeSync(tests []testCase, driver Driver) {
	fmt.Printf("Executing tests %+v sequentially\n", tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests in Sync
	for _, test := range tests {
		go LaunchTest(test, driver.portToListenAt, ch)

		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}

func main() {
	chromeDriver := Driver{"chromedriver", 9222, nil}
	chromeDriver.Start()
	defer chromeDriver.Stop()

	// Define tests here.
	tests := []testCase{
		// To define tests which do not use any proxy, pass an empty Proxy object like below.
		//{"http://www.httpvshttps.com/", Proxy{}},
		//{"https://www.httpvshttps.com/", Proxy{}},
		//
		//		// To define tests which use proxy, pass a Proxy object having the proxy's address
		{"http://www.httpvshttps.com/", Proxy{"http://localhost:8090"}},
		{"https://www.httpvshttps.com/", Proxy{"http://localhost:8090"}},
	}

	executeAsync(tests, chromeDriver)
	//executeSync(tests, chromeDriver)
}
