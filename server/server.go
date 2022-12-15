package server

import (
	"fmt"
	"net"
)

type Config struct {
	Name           string `yaml:"name"`
	Network        string `yaml:"network"`
	ListenAddress  net.IP `yaml:"listenAddress"`
	ListenPort     int    `yaml:"listenPort"`
	Ssl            bool   `yaml:"ssl"`
	BackendAddress []BackendAddress
}
type BackendAddress struct {
	Address          net.IP `yaml:"address"`
	Port             int    `yaml:"port"`
	NodeExporterPort int    `yaml:"nodeExporterPort"`
	HealthCheck      HealthCheck
}
type HealthCheck struct {
	Path     string `yaml:"path"`
	Interval int    `yaml:"interval"`
	Timeout  int    `yaml:"timeout"`
}

func NewServer(config *Config) {
	switch config.Network {
	case "tcp", "tcp4", "tcp6":
		TcpServer(config)
	case "udp", "udp4", "udp6":
		fmt.Println("currently udp is not supported")
	default:
		panic("Unknown network type")
	}

}
