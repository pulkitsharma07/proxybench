package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

//TODO: Add checks to instance variables of all the structs

type Driver struct {
	pathToExecutable string
	port             int
	process          *selenium.Service
}

func (d *Driver) Start() {

	service, err := selenium.NewChromeDriverService(d.pathToExecutable, d.port)
	if err != nil {
		log.Fatal(err)
	}
	d.process = service
}

func (d *Driver) Stop() {
	d.process.Stop()
}

type BrowserAutomator struct {
	// Port where chromedriver/geckodriver/Selenium JAR is listening
	ServicePort int
}

func (b *BrowserAutomator) StartChrome(proxyConf ...Proxy) selenium.WebDriver {
	caps := selenium.Capabilities{"browserName": "chrome"}

	if len(proxyConf) > 0 && proxyConf[0] != (Proxy{}) {
		chromeCaps := chrome.Capabilities{
			Args: []string{"--proxy-server=" + proxyConf[0].address},
		}
		caps.AddChrome(chromeCaps)
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", b.ServicePort))

	if err != nil {
		log.Fatal(err)
	}

	return wd
}
