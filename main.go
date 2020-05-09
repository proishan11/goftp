package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/proishan11/go-ftp/server"
)

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

		go server.HandleConnection(conn)
	}
}
