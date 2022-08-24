package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}

		io.WriteString(conn, "Hello from tcp server")
		conn.Close()
	}
}
