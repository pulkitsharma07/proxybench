package benchmark

func HeavyWebsites() *Websites {
	// Only one website for now, as Chromedriver crashes for some reason and return EOF for open url request.
	return NewWebsites("Heavy Websites", []string{"https://www.techcrunch.com"})
}
