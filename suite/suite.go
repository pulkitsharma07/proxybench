package suite

import (
	"fmt"
	"sync"

	"github.com/pulkitsharma07/proxybench/benchmark"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
)

type Suite struct {
	config  []config.Config
	tests   []benchmark.Benchmark
	results []result.Result
}

// Generates a new Suite, Generates a predefined list of  benchmarks to execute
// for each proxy.
func NewSimpleSuite(config []config.Config) *Suite {
	return &Suite{
		config,
		[]benchmark.Benchmark{
			benchmark.StressHTTP(),
			benchmark.StressHTTPS(),
			benchmark.HeavyWebsites(),
		},
		[]result.Result{},
	}
}

func executeWithConfig(config config.Config) {
}

func (t *Suite) Run() {
	fmt.Printf("Starting with Config: %v\n", t.config)
	var wg sync.WaitGroup

	for _, benchConfig := range t.config {
		wg.Add(1)
		go t.executeSync(benchConfig, &wg)
	}
	wg.Wait()
}

func (t *Suite) String() string {
	return fmt.Sprintf("ProxyConfig: %v, tests: %v \n", t.config, t.tests)
}

func (t *Suite) pushResults(res result.Result) {
	t.results = append(t.results, res)
}

func (t *Suite) Results() []result.Result {
	return t.results
}

func (t *Suite) executeSync(benchConfig config.Config, wg *sync.WaitGroup) {
	// Launch tests in Sync
	for _, test := range t.tests {
		fmt.Printf("\t\tStarting %+v\n", test)
		test.Run(benchConfig)
		t.pushResults(test.Results())
	}
	wg.Done()
}
