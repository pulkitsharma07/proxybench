package config

import "fmt"

type Proxy struct {
	Name    string
	Address string
}

func (p Proxy) String() string {
	return fmt.Sprintf("%s", p.Name)
}
