package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	httpRequest := "GET / HTTP/1.1\n" +
		"host: golang.org\n\n"
	conn, err := net.Dial("tcp", "golang.org:80") // port 80 - HTTP || port 443 - HTTPS
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	if _, err = conn.Write([]byte(httpRequest)); err != nil {
		fmt.Print(err)
		return
	}
	io.Copy(os.Stdout, conn)
	fmt.Println("Done")
}

/*
	func Dial(network, address string) (Conn, error)
	network - is protocol type as TCP, UDP, IP, unix, unixgram
	adress - net address like: 127.0.0.1 || 127.0.0.1:80 ||
	::1 || [2516:b7f0:3421:b16::71]:80

	returns
	interface net.Conn -> can be used as io.Reader & io.Writer

*/
