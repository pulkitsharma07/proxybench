package main

// TestCase generates testResult
type testCase interface {
	// Name specifies the testCase's Name. For Example: "HTTP Stress"
	Name() string

	// Runs the testcase using a WebDriver compatible server listening
	// on port <port> and pushed the result to the <done> channel on completion
	Run(config testConfig, port int, done chan testResult)
}
