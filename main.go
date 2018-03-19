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

type testCase struct {
	url string
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

	chrCaps := chrome.Capabilities{
		Args: []string{
			//"--proxy-server=http://localhost:8090",
		},
	}

	caps.AddChrome(chrCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal(err)
	}
	defer wd.Quit()

	if err := wd.Get(test.url); err != nil {
		panic(err)
	}

	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	timeToLoadStr, err := timeToLoad.Text()
	duration, err := strconv.ParseFloat(strings.Replace(timeToLoadStr, " s", "", -1), 64)

	done <- testResult{test, duration}
}

func main() {
	chromedriverPort := 9515
	chromedriver := startChromeDriver(chromedriverPort)
	defer chromedriver.Stop()

  tests := []testCase{
    testCase{"http://www.httpvshttps.com/"},
    testCase{"https://www.httpvshttps.com/"},
  }
	ch := make(chan testResult)

  for _, test := range tests {
    go LaunchTest(test, chromedriverPort, ch)
    time.Sleep(1 * time.Second)
  }

  for range tests {
    var res = <-ch
    fmt.Printf("Time for test %s : %f seconds\n", res.test.url, res.result)
  }
}
