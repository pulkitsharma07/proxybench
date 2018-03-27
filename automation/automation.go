package automation

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"

	"github.com/pulkitsharma07/proxybench/config"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

//TODO: Add checks to instance variables of all the structs

type Driver struct {
	pathToExecutable string
	Port             int
	process          *selenium.Service
}

// Starts chromedriver on random port.
func NewChromeDriver() *Driver {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		panic(err)
	}

	port, err := strconv.Atoi(strings.Split(listener.Addr().String(), ":")[1])
	if err != nil {
		panic(err)
	}
	listener.Close()

	return &Driver{
		"chromedriver",
		port,
		&selenium.Service{},
	}
}

func (d *Driver) Start() {
	service, err := selenium.NewChromeDriverService(d.pathToExecutable, d.Port)
	if err != nil {
		log.Fatal(err)
	}
	d.process = service
}

func (d *Driver) Stop() {
	d.process.Stop()
}

type BrowserAutomator struct {
	driver *Driver
	Wd     selenium.WebDriver
}

func NewChromeAutomator(proxyConf ...config.Proxy) *BrowserAutomator {
	chromedriver := NewChromeDriver()
	chromedriver.Start()

	caps := selenium.Capabilities{"browserName": "chrome"}

	if len(proxyConf) > 0 && proxyConf[0].Address != "" {
		chromeCaps := chrome.Capabilities{
			Args: []string{"--proxy-server=" + proxyConf[0].Address},
		}
		caps.AddChrome(chromeCaps)
	}

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", chromedriver.Port))

	if err != nil {
		log.Fatal(err)
	}

	return &BrowserAutomator{
		chromedriver,
		wd,
	}
}

func (b *BrowserAutomator) Stop() {
	b.Wd.Quit()
	b.driver.Stop()
}
