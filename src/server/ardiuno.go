package main;

import (
	"fmt"
	"net"
)

const UDP_MAX_PACKET_SIZE int = 2048

const serverTCPAddress string = ":8000"
const serverUDPAddress UDPAddr = net.ResolveUDPAddr("udp",":8000")
const teensyUDPAddress UDPAddr = net.ResolveUDPAddr("udp","192.168.1.100:8000")

func tcpSocket() {
	readBuf := make([]byte, 2048)
	
	listener, err := net.Listen("tcp", serverTCPAddress)
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected")
		bytesRead, err := conn.Read(readBuf)
		fmt.Println(bytesRead)
	}
}

func udpSocket() {
	udpConn, err := net.ListenUDP("udp", serverUDPAddress)
	if err != nil {
		panic(err)
	}
	for {
		buf := make([]byte, UDP_MAX_PACKET_SIZE)
	}
}

func startArduinoComs() {
	tcpSocket()
}
