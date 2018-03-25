package config

import "fmt"

type Config struct {
	ProxyToUse Proxy
}

func (c Config) String() string {
	return fmt.Sprintf("Proxy: [%s]", c.ProxyToUse)
}
