package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

// This example shows how to navigate to a http://play.golang.org page, input a
// short program, run it, and inspect its output.
//
// If you want to actually run this example:
//
//   1. Ensure the file paths at the top of the function are correct.
//   2. Remove the word "Example" from the comment at the bottom of the
//      function.
//   3. Run:
//      go test -test.run=Example$ github.com/tebeka/selenium
func main() {
	// Start a Selenium WebDriver server instance (if one is not already
	// running).
	const (
		// These paths will be different on your system.
		port = 9222
	)
	opts := []selenium.ServiceOption{}
	//selenium.SetDebug(true)
	service, err := selenium.NewChromeDriverService("chromedriver", port, opts...)
	if err != nil {
		log.Fatal(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer service.Stop()

	// Connect to the WebDriver instance running locally.
	caps := selenium.Capabilities{"browserName": "chrome"}

	chrCaps := chrome.Capabilities{
		Args: []string{
			"--proxy-server=http://localhost:8090",
		},
	}

	caps.AddChrome(chrCaps)
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		log.Fatal(err) // panic is used only as an example and is not otherwise recommended.
	}
	defer wd.Quit()

	// Navigate to the simple playground interface.
	if err := wd.Get("http://www.httpvshttps.com/"); err != nil {
		panic(err)
	}

	timeToLoad, err := wd.FindElement(selenium.ByCSSSelector, "#time")
	fmt.Print(timeToLoad.Text())

}
