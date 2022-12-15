package server

import (
	"fmt"
	"io"
	"net"
)

func proxy(backend string, c net.Conn) error {

	backendConn, err := net.Dial("tcp", backend)
	if err != nil {
		return fmt.Errorf("Error dialing backend: %s", err.Error())
	}
	//c -> bc
	go io.Copy(backendConn, c)
	//bc -> c
	go io.Copy(c, backendConn)
	return nil
}
