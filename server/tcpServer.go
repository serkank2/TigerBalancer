package server

import (
	"fmt"
	"log"
	"net"
)

func TcpServer(config *Config) {

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
