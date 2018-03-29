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

## Installation
* `go get github.com/pulkitsharma07/proxybench`

## Running the benchmark
* Make sure you have `chromedriver` present in your PATH.
* Start the desired proxies on different ports.
* Make sure you add and trust the certificates provided by the respective proxies, so that they can be considered as a trusted CA by the browser.
* Refer `proxybench.json.sample` and create `proxybench.json` similarly.
* Launch `proxybench`, it will look for `proxybench.json` in the current directory and run according to that config.

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
| Proxy: [charles]               | HTTP Stress    | [18.535000]            |
| Proxy: [charles]               | HTTPS Stress   | [6.322000]             |
| Proxy: [charles]               | Heavy Websites | [52.927620]            |
+--------------------------------+----------------+------------------------+
</pre>
* Generated using: browsermob v2.1.4, mitmproxy/mitmdump v3.0.3, charles 4.2.1
* `Direct` stands for the baseline performance, i.e. without using any proxy.

## Development
* Use `dep` to install dependencies
* Add proxy information in `proxybench.json`
* benchmarks are defined in `benchmarks/`
* You can define new benchmarks similarly, add them to `NewSimpleSuite` to execute.
* `go install`
* `proxybench` to run (Assuming GOBIN is added to your PATH)

## TODO
* Unit and integration tests.
* CLI option to skip the baseline (DIRECT) test.
* CLI options to launch tests in sync/async.
* Add wrapper around `Run` to measure the proxy's CPU/Memory usage (Need to take PID from user, or figure out from script?).Eventually render a timeseries graph.
* Randomize test order and most importantly do multiple test runs and report on standard deviation. (Thanks @mhills)
* Test reports, generate test reports confining to some format, each test case will have corresponding timeseries graphs of CPU, memory, disk, number of sockets etc of the proxy.
* Final plan is to give every proxy a score based on its performance in comparison to the baseline for the different benchmarks.
* Launch proxies from this script?


*PS: I made this mainly to learn `Go`. I will really appreciate if you can point out things in the code which I should improve upon. For example naming conventions, directory structure, or new benchmarks etc. Feel free to open any issues.*

**Thank You !**
