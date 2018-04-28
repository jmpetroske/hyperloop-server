package main;

import (
    "fmt"
    "net"
	"log"
)

type DataFrame struct {
    speed int
    // TODO
}

const UDP_MAX_PACKET_SIZE int = 2048

const serverTCPAddress string = ":8000"
var serverUDPAddress, _ = net.ResolveUDPAddr("udp",":8000")
var teensyUDPAddress, _ = net.ResolveUDPAddr("udp","192.168.1.100:8000")

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
		// TODO check that this is a connection to the teensy
        log.Println("Got TCP connection with teensy")
        bytesRead, err := conn.Read(readBuf)
        log.Println(bytesRead)
    }
}

func udpSocket() {
    udpConn, err := net.ListenUDP("udp", serverUDPAddress)
    if err != nil {
        panic(err)
    }
    buf := make([]byte, UDP_MAX_PACKET_SIZE)
    for {
        n, senderAddr, err := udpConn.ReadFromUDP(buf)
        if err != nil {
            log.Println(err)
            continue
        }
        if !senderAddr.IP.Equal(teensyUDPAddress.IP) ||
            senderAddr.Port != teensyUDPAddress.Port {
            log.Println("Got UDP packet from non teensy address")
            continue
        }

        log.Println("Got UDP packet from teensy: " + string(buf[0:n]))
    }
}

func parseDataPacket(dataPacket []byte) *DataFrame {
    return &DataFrame{}
}

func startArduinoComs() {
    go udpSocket()
    tcpSocket()
}
