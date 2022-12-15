package main

import (
	"fmt"
	"net"
	"tigerblancer/loader"
	"tigerblancer/server"
)

func main() {

	config := loader.ConfigLoad()
	for _, v := range config.Settings {

		Status := true
		ListenAddress := net.ParseIP(v.ListenAddress)

		switch v.Network {
		case "tcp", "tcp4", "tcp6", "udp", "udp4", "udp6":
		default:
			fmt.Println("Network is not valid")
			Status = false
		}

		if ListenAddress == nil {
			fmt.Println("ListenAddress is not valid")
			Status = false
		}

		if Status {
			fmt.Println("Check is OK")
		}

	}

	app := server.NewServer(&server.Config{
		Name:          "test",
		Network:       "tcp",
		ListenAddress: "192.168.1.1",
		ListenPort:    8080,
		Ssl:           false,
		BackendAddress: &server.BackendAddress{
			Address:          "192.168.1.1",
			Port:             8080,
			NodeExporterPort: 8080,
			HealthCheck: &server.HealthCheck{
				Path:     "/health",
				Interval: 60,
				Timeout:  30,
			},
		},
	})
	fmt.Println(app)
}
