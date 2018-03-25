package main

import (
	"fmt"

  "github.com/pulkitsharma07/proxybench/launcher"
  "github.com/pulkitsharma07/proxybench/config"
)

func main() {
  suite := launcher.Launch([]config.Config{{config.Proxy{}}})
  fmt.Printf("%+v", suite.Results())
}
