package main

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

// HTTPTestCase implements testCase
type HTTPSTestCase struct {
	name string
}

func (t HTTPSTestCase) String() string {
	return t.name
}

func (t HTTPSTestCase) Run(config testConfig, port int, done chan testResult) {
	b := &BrowserAutomator{port}
	wd := b.StartChrome(config.proxyConfig)

	if err := wd.Get("https://www.httpvshttps.com"); err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)
	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	if err != nil {
		log.Fatal("Failed to parse results")
	}

	timeToLoadStr, err := timeToLoad.Text()
	duration, err := strconv.ParseFloat(strings.Replace(timeToLoadStr, " s", "", -1), 64)

	wd.Quit()
	done <- NewLoadTimeResult(t, "PageLoadTime(sec)", duration)
}
