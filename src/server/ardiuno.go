package main

import (
	"encoding/binary"
	"io"
	"log"
	"math"
	"net"
	"reflect"
)

const SEND_RATIO = 100 // Send 1 in 100 packets
const UDP_MAX_PACKET_SIZE int = 2048

const serverAddress string = ":8888"

var serverUDPAddress, _ = net.ResolveUDPAddr("udp", serverAddress)
var teensyUDPAddress, _ = net.ResolveUDPAddr("udp", "192.168.1.100:8888")

func debugTcpSocket() {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(err)
	}
	
	log.Println("Wating for connection with teensy")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	log.Println("Got TCP connection with teensy for testing")

	// bytesRead, err := conn.Read(readBuf)
	// log.Println(bytesRead)

	go func() {
		for {
			_, err := conn.Write((<-commandChan).WriteCommand())
			log.Println("Got a command in arduino coms, sending to teensy")
			if err != nil {
				log.Print("Parsing error: ")
				log.Println(err)
				log.Println("Closing connection with teensy")
				conn.Close()
				return
			}
		}
	}()

	for {
		dp, err := tcpDataPacketParser(conn)
		if err != nil {
			log.Print("Parsing error: ")
			log.Println(err)
			log.Println("Closing connection with teensy")
			conn.Close()
			break
		}
		log.Println("Got a packet")
		latestDataMutex.Lock()
		latestData = dp
		latestDataMutex.Unlock()
	}
}

func tcpSocket() {
	log.Println("ERROR: don't use this mode")
	// // readBuf := make([]byte, 2048)
	// listener, err := net.Listen("tcp", serverAddress)
	// if err != nil {
	// 	panic(err)
	// }
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	// TODO check that this is a connection to the teensy
	// 	log.Println("Got TCP connection with teensy")
	// 	// bytesRead, err := conn.Read(readBuf)
	// 	// log.Println(bytesRead)
	// 	for {
	// 		_, err := conn.Write((<-commandChan).WriteCommand())
	// 		if err != nil {
	// 			log.Println(err)
	// 		}
	// 	}
	// }
}

func udpSocket() {
	log.Println("ERROR: don't use this mode")
	// udpConn, err := net.ListenUDP("udp", serverUDPAddress)
	// if err != nil {
	// 	panic(err)
	// }

	// go func() {
	// 	<-abortChan
	// 	for {
	// 		udpConn.WriteToUdp([] TODO, teensyUDPAddress)
	// 	}
	// }()

	// buf := make([]byte, UDP_MAX_PACKET_SIZE)
	// for {
	// 	n, senderAddr, err := udpConn.ReadFromUDP(buf)
	// 	if err != nil {
	// 		log.Println(err)
	// 		continue
	// 	}
	// 	if !senderAddr.IP.Equal(teensyUDPAddress.IP) ||
	// 		senderAddr.Port != teensyUDPAddress.Port {
	// 		log.Println("Got UDP packet from non teensy address")
	// 		continue
	// 	}

	// 	log.Println("Got UDP packet from teensy: " + string(buf[0:n]))
	// 	dataPacket := parseDataPacket(buf[0:n])
	// 	logDataPacket(dataPacket)
	// 	latestDataMutex.Lock()
	// 	latestData = dataPacket
	// 	latestDataMutex.Unlock()
	// }
}

func tcpDataPacketParser(conn net.Conn) (*DataPacket, error) {
	retval := DataPacket{}
	reflectValue := reflect.ValueOf(&retval).Elem()
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)

		switch field.Kind() {
		case reflect.Uint32:
			b := make([]byte, 4)
			_, err := io.ReadFull(conn, b)
			if err != nil {
				return nil, err
			}
			field.SetUint(uint64(binary.LittleEndian.Uint32(b)))
		case reflect.Float32:
			b := make([]byte, 4)
			_, err := io.ReadFull(conn, b)
			if err != nil {
				return nil, err
			}
			field.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
		case reflect.Bool:
			b := make([]byte, 1)
			_, err := io.ReadFull(conn, b)
			if err != nil {
				return nil, err
			}
			field.SetBool(b[0] != 0)
		default:
			log.Fatal("Error parsing data from teensy, bad use of reflection")
		}
	}

	return &retval, nil
}

func startArduinoComs() {
	go udpSocket()
	tcpSocket()
}

func startDebugArduinoComs() {
	debugTcpSocket()
}
