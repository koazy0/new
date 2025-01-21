package config

import "fmt"

type ES struct {
	IP       string `yaml:"ip"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

func (e ES) URL() string {
	return fmt.Sprintf("http://%s:%d", e.IP, e.Port)
}
