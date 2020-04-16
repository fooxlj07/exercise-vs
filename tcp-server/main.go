package main

import (
	"log"
	"net"
)

const (
	connType = "tcp"
	connHost = "localhost"
	connPort = "8080"
)

func main() {
	listener, err := net.Listen(connType, connHost+":"+connPort)
	if err != nil {
		log.Fatal("tcp server listener error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("fail to get connectiong")
		}
		go func(conn net.Conn) {
			connHandler := NewConnectionHandler(conn)
			connHandler.start()
		}(conn)
	}
}
