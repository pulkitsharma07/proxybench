package config

import "fmt"

type Proxy struct {
	Name    string
	Address string
}

func (p Proxy) String() string {
	return fmt.Sprintf("Name: %s, Address: %s", p.Name, p.Address)
}
