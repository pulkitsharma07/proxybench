package config

type Proxy struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func (p Proxy) String() string {
	return fmt.Sprintf("%s", p.Name)
}
