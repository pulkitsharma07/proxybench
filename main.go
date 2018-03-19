package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Proxy struct {
	address string
}

type testCase struct {
	url   string
	proxy Proxy
}

type testResult struct {
	test   testCase
	result float64 //TODO: Make this generic
}

func startChromeDriver(port int) *selenium.Service {
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService("chromedriver", port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return service
}

func LaunchTest(test testCase, port int, done chan testResult) {
	caps := selenium.Capabilities{"browserName": "chrome"}

	// Use proxy if testCase.proxy is set
	if test.proxy != (Proxy{}) {
		chromeCaps := chrome.Capabilities{
			Args: []string{"--proxy-server=" + test.proxy.address},
		}
		caps.AddChrome(chromeCaps)
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal(err)
	}

	if err := wd.Get(test.url); err != nil {
		panic(err)
	}

	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	timeToLoadStr, err := timeToLoad.Text()
	duration, err := strconv.ParseFloat(strings.Replace(timeToLoadStr, " s", "", -1), 64)

	wd.Quit()
	done <- testResult{test, duration}
}

func executeAsync(tests []testCase, chromeDriverPort int) {
	fmt.Printf("Executing tests %+v parallely\n", tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests parallely
	for _, test := range tests {
		go LaunchTest(test, chromeDriverPort, ch)

		// Do not fire concurrent start requests to chromedriver.
		time.Sleep(1 * time.Second)
	}

	for range tests {
		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}

func executeSync(tests []testCase, chromeDriverPort int) {
	fmt.Printf("Executing tests %+v sequentially\n", tests)

	// channel for gathering results
	ch := make(chan testResult)

	// Launch tests in Sync
	for _, test := range tests {
		go LaunchTest(test, chromeDriverPort, ch)

		var res = <-ch
		fmt.Printf("TestResult %+v \n", res)
	}
}

func main() {

	// Start chromedriver
	chromedriverPort := 9515
	chromedriver := startChromeDriver(chromedriverPort)

	// Define tests here.
	tests := []testCase{
		// To define tests which do not use any proxy, pass an empty Proxy object like below.
		{"http://www.httpvshttps.com/", Proxy{}},
		{"https://www.httpvshttps.com/", Proxy{}},

		// To define tests which use proxy, pass a Proxy object having the proxy's address
		{"http://www.httpvshttps.com/", Proxy{"http://localhost:8090"}},
		{"https://www.httpvshttps.com/", Proxy{"http://localhost:8090"}},
	}

	executeAsync(tests, chromedriverPort)
	//executeSync(tests, chromedriverPort)

	chromedriver.Stop()
}
