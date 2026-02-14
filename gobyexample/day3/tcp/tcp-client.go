package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	port := ":2121"

	conn, err := net.Dial("tcp", "127.0.0.1"+port)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	io.Copy(os.Stdout, conn)
	fmt.Println("\nDone")

}
