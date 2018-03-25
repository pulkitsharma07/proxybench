package benchmark

import (
	"github.com/pulkitsharma07/proxybench/config"
	"github.com/pulkitsharma07/proxybench/result"
)

// TestCase generates testResult
type Benchmark interface {
	// Gets the Benchmark's Name For Example: "HTTP Stress"
	String() string

	// Runs the benchmark using a WebDriver compatible server listening
	// on port <port> and pushes the result to the <done> channel on completion
	Run(config config.Config, port int)

	// Should return the result(s) generated for this benchmark
	Results() result.Result
}
