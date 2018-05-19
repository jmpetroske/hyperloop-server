package main

import (
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"math"
	"net"
	"reflect"
)

const SEND_RATIO = 100 // Send 1 in 100 packets
const UDP_MAX_PACKET_SIZE int = 2048

const serverTCPAddress string = ":8888"

var serverUDPAddress, _ = net.ResolveUDPAddr("udp", ":8888")
var teensyUDPAddress, _ = net.ResolveUDPAddr("udp", "192.168.1.100:8888")

func debugTcpSocket() {
	listener, err := net.Listen("tcp", serverTCPAddress)
	if err != nil {
		panic(err)
	}
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
			log.Println("Got a command")
			if err != nil {
				log.Println(err)
			}
		}
	}()

	for {
		dp, err := tcpDataPacketParser(conn)
		if err != nil {
			log.Println(err)
			continue
		}
		logDataPacket(dp)
		latestDataMutex.Lock()
		latestData = dp
		latestDataMutex.Unlock()
		// log.Print("Got a data packet")
		// log.Println(dp)
	}
}

func tcpSocket() {
	// readBuf := make([]byte, 2048)
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
			_, err := conn.Write((<-commandChan).WriteCommand())
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
	dataVals := bytes.Split(data, []byte{','})
	retval := DataPacket{}
	reflectValue := reflect.ValueOf(&retval).Elem()
	for i := 0; i < reflectValue.NumField(); i++ {
		field := reflectValue.Field(i)

		switch field.Kind() {
		case reflect.Uint32:
			field.SetUint(uint64(binary.LittleEndian.Uint32(dataVals[i])))
		case reflect.Float32:
			field.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(dataVals[i]))))
		case reflect.Bool:
			noZeros := true
			for _, b := range dataVals[i] {
				noZeros = noZeros || b == 0
			}
			field.SetBool(noZeros)
		default:
			log.Println("Error parsing data from teensy, bad use of reflection")
		}
	}

	return &retval
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
				log.Print("Parsing error: ")
				log.Println(err)
				return nil, err
			}
			field.SetUint(uint64(binary.LittleEndian.Uint32(b)))
		case reflect.Float32:
			b := make([]byte, 4)
			_, err := io.ReadFull(conn, b)
			if err != nil {
				log.Print("Parsing error: ")
				log.Println(err)
				return nil, err
			}
			field.SetFloat(float64(math.Float32frombits(binary.LittleEndian.Uint32(b))))
		case reflect.Bool:
			b := make([]byte, 1)
			_, err := io.ReadFull(conn, b)
			if err != nil {
				log.Print("Parsing error: ")
				log.Println(err)
				return nil, err
			}
			field.SetBool(b[0] != 0)
		default:
			log.Println("Error parsing data from teensy, bad use of reflection")
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
