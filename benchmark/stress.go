package benchmark

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/pulkitsharma07/proxybench/automation"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
	"github.com/tebeka/selenium"
)

// Stress implements Benchmark
type Stress struct {
	name   string
	url    string
	result result.LoadTimeResult
}

func NewStress(name string, https bool) *Stress {
	if https {
		return &Stress{
			name,
			"https://www.httpvshttps.com",
			result.LoadTimeResult{},
		}
	} else {
		return &Stress{
			name,
			"http://www.httpvshttps.com",
			result.LoadTimeResult{},
		}
	}
}

func (t *Stress) String() string {
	return t.name
}

func (t *Stress) Results() result.Result {
	return t.result
}

func (t *Stress) Run(config config.Config) {
	chrome := automation.NewChromeAutomator(config.ProxyToUse)
	defer chrome.Stop()
	wd := chrome.Wd

	if err := wd.Get(t.url); err != nil {
		panic(err)
	}

	//               Fetch the elapsed time from the DOM.
	// Wait just to ensure that page has completely loaded and timer has stopped.
	time.Sleep(2 * time.Second)

	// Get the timer DOM element
	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	if err != nil {
		fmt.Printf("Unable to find score..\n")
		t.result = result.NewLoadTimeResult(config, t.name+" (FAILED)", "PageLoadTime(sec)", -1.0)
		return
	}

	// Parse score to float (Will help in computing other things in future)
	timeToLoadStr, err := timeToLoad.Text()
	duration, err := strconv.ParseFloat(strings.Replace(timeToLoadStr, " s", "", -1), 64)

	t.result = result.NewLoadTimeResult(config, t.name, "PageLoadTime(sec)", duration)
}
