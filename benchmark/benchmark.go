package benchmark

import (
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
)

// TestCase generates testResult
type Benchmark interface {
	// Gets the Benchmark's Name For Example: "HTTP Stress"
	String() string

	// Runs the benchmark using by launching a WebDriver compatible server listening
	// on a random port.
	Run(config config.Config)

	// Should return the result(s) generated for this benchmark
	Results() result.Result
}
