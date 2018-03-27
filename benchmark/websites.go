package benchmark

import (
	"fmt"
	"time"

	"github.com/pulkitsharma07/proxybench/automation"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
)

// Websites implements Benchmark
type Websites struct {
	name   string
	urls   []string
	result result.LoadTimeResult
}

func NewWebsites(name string, urls []string) *Websites {
	return &Websites{
		name,
		urls,
		result.LoadTimeResult{},
	}
}

func (t *Websites) String() string {
	return t.name
}

func (t *Websites) Results() result.Result {
	return t.result
}

func (t *Websites) Run(config config.Config) {
	chrome := automation.NewChromeAutomator(config.ProxyToUse)
	defer chrome.Stop()
	wd := chrome.Wd

	start := time.Now()

	for _, url := range t.urls {
		if err := wd.Get(url); err != nil {
			fmt.Printf("Failed because of: %+v\n", err)
			elapsed := time.Now().Sub(start)
			t.result = result.NewLoadTimeResult(config, t.name, "PageLoadTime(sec)", elapsed.Seconds())
			return
		}
	}
	elapsed := time.Now().Sub(start)
	t.result = result.NewLoadTimeResult(config, t.name, "PageLoadTime(sec)", elapsed.Seconds())
}
