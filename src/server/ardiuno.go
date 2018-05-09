package main

import (
	"log"
	"net"
	"sync"
)

const SEND_RATIO = 100 // Send 1 in 100 packets
const UDP_MAX_PACKET_SIZE int = 2048

const serverTCPAddress string = ":8000"

var serverUDPAddress, _ = net.ResolveUDPAddr("udp", ":8000")
var teensyUDPAddress, _ = net.ResolveUDPAddr("udp", "192.168.1.100:8000")

var latestDataMutex = &sync.Mutex{}
var latestData *DataPacket

var commandChan = make(chan PhotonCommand, 3)

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
		// bytesRead, err := conn.Read(readBuf)
		// log.Println(bytesRead)
		for {
			bytesRead, err := conn.Write((<-commandchan).WriteCommand())
			if err != nil {
				log.Println(err)
			}
		}
	}
}

func udpSocket() {
	udpConn, err := net.ListenUDP("udp", serverUDPAddress)
	if err != nil {
		panic(err)
	}
	buf := make([]byte, UDP_MAX_PACKET_SIZE)
	numSent := 0
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
		dataPacket := parseDataPacket(buf[0:n])
		logDataPacket(dataPacket)
		latestDataMutex.Lock()
		latestData = dataPacket
		latestDataMutex.Unlock()

		numSent++
	}
}

func parseDataPacket(data []byte) *DataPacket {
	return &DataPacket{}
}

func startArduinoComs() {
	go udpSocket()
	go tcpSocket()
}
