package server

import (
	"fmt"
	"io"
	"log"
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

func (config *Config) Server() {

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: config.ListenAddress, Port: config.ListenPort})
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting: %s", err.Error())
		}
		backend := config.BackendAddress[chooseBackend()].Address.String() + ":" + fmt.Sprint(config.BackendAddress[chooseBackend()].Port)

		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("Error proxying: %s", err.Error())
			}
		}()
	}

}

func NewServer(config *Config) {
	config.Server()
}

func proxy(backend string, c net.Conn) error {

	bc, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("Error dialing backend: %s", err.Error())
	}
	//c -> bc
	go io.Copy(bc, c)
	//bc -> c
	go io.Copy(c, bc)
	return nil
}

func chooseBackend() int {
	return 0
}
