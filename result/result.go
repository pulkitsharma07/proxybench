package result

import (
	"fmt"

	"github.com/pulkitsharma07/proxybench/config"
)

// testResultField has the actual result for example:
// "duration": 2.33
type Field struct {
	Name  string
	Value float64 // TODO: make this generic
}

func (f Field) String() string {
	return fmt.Sprintf("%f", f.Value)
}

// Result has many Fiedls
type Result interface {
	Results() []Field

	// Should pretty print the results
	String() string

	// Should provide the name of the benchmark it is the result of.
	Benchmark() string

	// Should provide the Config with which this benchmark was run
	Config() config.Config
}

//LoadTimeResult is a Result
type LoadTimeResult struct {
	config    config.Config
	benchmark string
	results   []Field
}

func NewLoadTimeResult(config config.Config, benchmark string, name string, duration float64) LoadTimeResult {
	return LoadTimeResult{
		config,
		benchmark,
		[]Field{{
			name,
			duration,
		}},
	}
}

func (l LoadTimeResult) Results() []Field {
	return l.results
}

func (l LoadTimeResult) Benchmark() string {
	return l.benchmark
}

func (l LoadTimeResult) Config() config.Config {
	return l.config
}

func (l LoadTimeResult) String() string {
	return fmt.Sprintf("Config: %v, Benchmark: %v, Result: %v", l.config, l.benchmark, l.results)
}
