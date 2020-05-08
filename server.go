package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const (
	defaultPort = 8000
	defaultURL  = "localhost"
)

func handleConn(c net.Conn) {
	io.Copy(c, c)
	c.Close()
	return
}

func main() {
	var port string

	if len(os.Args) == 1 {
		port = "8000"
	} else {
		port = os.Args[1]
	}

	listener, err := net.Listen("tcp", "localhost:"+port)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Server started on port %s \n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(conn)
	}
}
