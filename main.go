package main

import "fmt"

func main() {
	config := ConfigLoad()
	fmt.Println(config.ListenerProtocol)
}
