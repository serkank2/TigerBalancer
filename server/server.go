package server

import (
	"fmt"
	"io"
	"log"
	"net"
)

var (
	listenAddr = "127.0.0.1:8080"
	server     = []string{
		"localhost:8081",
		"localhost:8082",
		"localhost:8083",
	}
)

func Server() {

	listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.IP{127, 0, 0, 1}, Port: 8080})
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting: %s", err.Error())
		}
		backend := chooseBackend()
		go func() {
			err := proxy(backend, conn)
			if err != nil {
				log.Printf("Error proxying: %s", err.Error())
			}
		}()

	}

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

func chooseBackend() string {
	return server[0]
}
