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
		fmt.Println(v.Name)
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

		var BackendAddressList []server.BackendAddress

		for _, v := range v.BackendAddress {
			BackendAddressList = append(BackendAddressList, server.BackendAddress{
				Address:          net.ParseIP(v.Address),
				Port:             v.Port,
				NodeExporterPort: v.NodeExporterPort,
				HealthCheck: server.HealthCheck{
					Path:     v.HealthCheck.Path,
					Interval: v.HealthCheck.Interval,
					Timeout:  v.HealthCheck.Timeout,
				},
			})
			if BackendAddressList == nil {
				fmt.Println("BackendAddressList is not valid")
			}
		}
		if Status {
			go server.NewServer(&server.Config{
				Name:           v.Name,
				Network:        v.Network,
				ListenAddress:  net.ParseIP(v.ListenAddress),
				ListenPort:     v.ListenPort,
				Ssl:            v.Ssl,
				BackendAddress: BackendAddressList,
			})
		}
	}
	for {

	}
}
