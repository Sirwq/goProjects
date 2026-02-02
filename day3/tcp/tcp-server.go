package main

import (
	"fmt"
	"net"
)

func main() {
	message := "Hello! Hi! Hola! Привет!"
	port := ":2121"
	listener, err := net.Listen("tcp", port)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	fmt.Printf("Server listening on port: %s", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write([]byte(message))
		conn.Close()
	}
}
