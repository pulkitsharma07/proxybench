## proxybench
Benchmark different proxies on basis of the following:

| test                    |   Implemented       | Description |
|:------------------------|:-------------------:|:------------|
|HTTP Stress              | :white_check_mark:  | Goes to http://httpvshttps.com and fetches the time to load everything.|
|HTTPS Stress (HTTP/2)    | :white_check_mark:  | Goes to https://httpvshttps.com and fetches the time to load everything.|
|HTTPS Stress (HTTP/1.1)  |                     ||
|WebSocket                |                     ||
|Video Streaming          |                     ||
|WebSocket Secure         |                     ||
|\<TO ADD MORE\>          |                     ||

## Reports
Generates the following:
<pre>
+--------------------------------+--------------+------------------------+
|          PROXY CONFIG          |  BENCHMARK   | COMPLETED IN (SECONDS) |
+--------------------------------+--------------+------------------------+
| Proxy: [Direct]                | HTTPS Stress | [4.592000]             |
| Proxy: [Direct]                | HTTP Stress  | [15.988000]            |
| Proxy: [browsermob (Legacy)]   | HTTPS Stress | [15.855000]            |
| Proxy: [browsermob (Legacy)]   | HTTP Stress  | [17.401000]            |
| Proxy: [browsermob             | HTTPS Stress | [17.774000]            |
| (LittleProxy)]                 |              |                        |
| Proxy: [browsermob             | HTTP Stress  | [16.448000]            |
| (LittleProxy)]                 |              |                        |
| Proxy: [mitmproxy]             | HTTPS Stress | [44.268000]            |
| Proxy: [mitmproxy]             | HTTP Stress  | [32.880000]            |
+--------------------------------+--------------+------------------------+
</pre>
Generated using:
browsermob v2.1.4, mitmproxy v2.0.2


## development
* Use `dep` to install dependencies
* Add proxy information in `proxybench.json`
* benchmarks are defined in `benchmarks/`
* You can define new benchmarks similarly, add them to `NewSimpleSuite` to execute.
* `go run *go` to run.

### TODO
* Unit and integration tests.
* CLI options to launch tests in sync/async
* Add wrapper around `Run` to measure the proxy's CPU/Memory usage (Need to take PID from user, or figure out from script?).Eventually render a timeseries graph.
* Test reports, generate test reports confining to some format, each test case will have corresponding timeseries graphs of CPU, memory, disk, number of sockets etc of the proxy.
* Launch proxies from this script?
