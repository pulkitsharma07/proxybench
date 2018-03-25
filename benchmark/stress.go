package benchmark

import (
	"log"
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

func (t *Stress) Run(config config.Config, port int) {
	b := automation.BrowserAutomator{port}
	wd := b.StartChrome(config.ProxyToUse)
	defer wd.Quit()

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

	t.result = result.NewLoadTimeResult(config, t.name, "PageLoadTime(sec)", duration)
}
