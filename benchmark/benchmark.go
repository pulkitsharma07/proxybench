package benchmark

// TestCase generates testResult
type Benchmark interface {
	// Gets the Benchmark's Name For Example: "HTTP Stress"
	String() string

	// Runs the benchmark using a WebDriver compatible server listening
	// on port <port> and pushes the result to the <done> channel on completion
	Run(config testConfig, port int, done chan testResult)
}
