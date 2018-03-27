## proxybench
Benchmark different proxies on basis of the following:

| test                    |   Implemented       | Description |
|:------------------------|:-------------------:|:------------|
|HTTP Stress              | :white_check_mark:  | records the time it takes to finish the test on http://httpvshttps.com  |
|HTTPS Stress (HTTP/2)    | :white_check_mark:  | records the time it takes to finish the test on https://httpvshttps.com |
| Heavy Websites          | :white_check_mark:  | records the time it takes to open websites like https://techcrunch.com/ |
|HTTPS Stress (HTTP/1.1)  |    ||
|WS Support               |                     ||
|WSS Support         |                     ||
|\<TO ADD MORE\>          |                     ||

## Reporting
Generates the following:
<pre>
+--------------------------------+----------------+------------------------+
|          PROXY CONFIG          |   BENCHMARK    | COMPLETED IN (SECONDS) |
+--------------------------------+----------------+------------------------+
| Proxy: [Direct]                | HTTP Stress    | [15.575000]            |
| Proxy: [Direct]                | HTTPS Stress   | [2.757000]             |
| Proxy: [Direct]                | Heavy Websites | [20.703525]            |
| Proxy:                         | HTTP Stress    | [17.920000]            |
| [browsermob(littleproxy)]      |                |                        |
| Proxy:                         | HTTPS Stress   | [17.053000]            |
| [browsermob(littleproxy)]      |                |                        |
| Proxy:                         | Heavy Websites | [40.902084]            |
| [browsermob(littleproxy)]      |                |                        |
| Proxy: [browsermob(legacy)]    | HTTP Stress    | [20.777000]            |
| Proxy: [browsermob(legacy)]    | HTTPS Stress   | [17.399000]            |
| Proxy: [browsermob(legacy)]    | Heavy Websites | [31.580184]            |
| Proxy: [mitmproxy]             | HTTP Stress    | [39.376000]            |
| Proxy: [mitmproxy]             | HTTPS Stress   | [47.631000]            |
| Proxy: [mitmproxy]             | Heavy Websites | [198.824246]           |
| Proxy: [mitmdump]              | HTTP Stress    | [17.960000]            |
| Proxy: [mitmdump]              | HTTPS Stress   | [8.147000]             |
| Proxy: [mitmdump]              | Heavy Websites | [41.605729]            |
+--------------------------------+----------------+------------------------+
</pre>
* Generated using: browsermob v2.1.4, mitmproxy/mitmdump v3.0.3
* `Direct` stands for the baseline performance, i.e. without using any proxy.

## Dependencies
* Make sure you have `chromedriver` present in your PATH.

## Development
* Use `dep` to install dependencies
* Add proxy information in `proxybench.json`
* benchmarks are defined in `benchmarks/`
* You can define new benchmarks similarly, add them to `NewSimpleSuite` to execute.
* `go install`
* `proxybench` to run (Assuming GOBIN is added to your PATH)

### TODO
* Unit and integration tests.
* CLI options to launch tests in sync/async
* Add wrapper around `Run` to measure the proxy's CPU/Memory usage (Need to take PID from user, or figure out from script?).Eventually render a timeseries graph.
* Test reports, generate test reports confining to some format, each test case will have corresponding timeseries graphs of CPU, memory, disk, number of sockets etc of the proxy.
* Final plan is to give every proxy a score based on its performance in comparison to the baseline for the different benchmarks.
* Launch proxies from this script?


*PS: I made this mainly to learn `Go`. I will really appreciate if you can point out things in the code which I should improve upon. For example naming conventions, directory structure, or new benchmarks etc. Feel free to open any issues.*

**Thank You !**
