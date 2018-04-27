package main;

import (
	"fmt"
	"net"
)

func startArduinoComs() {
	readBuf := make([]byte, 2048)
	
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		// handle error
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected")
		fmt.Println("%b", conn)
		bytesRead, err := conn.Read(readBuf)
		fmt.Println(bytesRead)
	}
}
