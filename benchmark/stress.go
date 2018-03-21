package benchmark

import (
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

// Stress implements Benchmark
type Stress struct {
	name string
	url  string
}

func NewStress(name string, https bool) Stress {
	if https {
		return Stress{
			name,
			"https://www.httpvshttps.com",
		}
	} else {
		return Stress{
			name,
			"http://www.httpvshttps.com",
		}
	}
}

func (t Stress) String() string {
	return t.name
}

func (t Stress) Run(config testConfig, port int, done chan testResult) {
	b := &BrowserAutomator{port}
	wd := b.StartChrome(config.proxyConfig)

	if err := wd.Get(t.url); err != nil {
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
