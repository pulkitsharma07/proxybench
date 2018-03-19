## proxybench
Benchmark different proxies on the basis of different metrics.

## development
* Get the Go selenium bindings: `go get github.com/tebeka/selenium`
* Currently HTTP/HTTPS tests are defined in `main.go`. Launch different proxies on different ports and add respective ports in the `tests` array.
You can then uncomment `executeSync` / `executeAsync`.
* `go run main.go` to launch tests.

### TODO
* Create a JSON config file for defining tests.
* Unit and integration tests.
* CLI options to launch tests in sync/async
* Need to define more types of `testCase`s, the currently defined are specific for HTTP vs HTTPS comparison.
* Add wrapper around `LaunchTest` to measure the proxy's CPU/Memory usage (Need to take PID from user, or figure out from script?).Eventually render a timeseries graph.
* Test reports, generate test reports confining to some format, each test case will have corresponding timeseries graphs of CPU, memory, disk, number of sockets etc of the proxy.
* Launch proxies from this script?
