package main

import (
	"fmt"
	"log"
	"net"
)

const listenAddr = "localhost:4000"

func main() {
	l, err := net.Listen("tcp", listenAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Listening on", listenAddr)
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Fprintln(conn, "Hello, world")
		conn.Close()
	}
}
