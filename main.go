package main

import (
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Printf("err %v", err)
	}

	hub := newHub()
	go hub.run()
}
