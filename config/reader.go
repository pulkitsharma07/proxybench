package config

import (
	"encoding/json"
	"os"
)

type ConfigJson struct {
	Proxies []Proxy `json:"proxies"`
}

func ReadConfig() []Config {
	fileReader, err := os.Open("proxybench.json")
	if err != nil {
		panic("Cannot find proxybench.json")
	}

	var configJson ConfigJson
	json.NewDecoder(fileReader).Decode(&configJson)

	suiteConfig := []Config{{Proxy{"Direct", ""}}}

	for _, proxy := range configJson.Proxies {
		suiteConfig = append(suiteConfig, Config{proxy})
	}

	return suiteConfig
}
