package suite

import (
	"fmt"

	"github.com/pulkitsharma07/proxybench/benchmark"
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
)

type Suite struct {
	config        []config.Config
	tests         []benchmark.Benchmark
	webDriverPort int
	results       []result.Result
}

// Generates a new Suite, Generates a predefined list of  benchmarks to execute
// for each proxy.
func NewSimpleSuite(config []config.Config, webDriverPort int) *Suite {
	return &Suite{
		config,
		[]benchmark.Benchmark{
			benchmark.StressHTTPS(),
			benchmark.StressHTTP(),
		},
		webDriverPort,
		[]result.Result{},
	}
}

func (t *Suite) Run() {
	fmt.Printf("Starting with Config: %v\n", t.config)
	for _, benchConfig := range t.config {
		fmt.Printf("\tLaunching benchmark(s) for: %+v\n", benchConfig)
		t.executeSync(benchConfig)
	}
}

func (t *Suite) String() string {
	return fmt.Sprintf("ProxyConfig: %v, Port: %v, tests: %v \n", t.config, t.webDriverPort, t.tests)
}

func (t *Suite) pushResults(res result.Result) {
	t.results = append(t.results, res)
}

func (t *Suite) Results() []result.Result {
	return t.results
}

func (t *Suite) executeSync(benchConfig config.Config) {
	// Launch tests in Sync
	for _, test := range t.tests {
		fmt.Printf("\t\tStarting %+v\n", test)
		test.Run(benchConfig, t.webDriverPort)
		t.pushResults(test.Results())
	}
}
