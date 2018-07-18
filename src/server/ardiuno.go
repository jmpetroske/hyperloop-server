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

// readData: if true, read DataPackets from the TCP connection
func tcpComs(readData bool) {
	listener, err := net.Listen("tcp", serverAddress)
	if err != nil {
		panic(err)
	}

	log.Println("Waiting for connection with teensy")
	conn, err := listener.Accept()
	if err != nil {
		panic(err)
	}
	log.Println("Got TCP connection with teensy")

	if readData {
		// goroutine that reads DataPackets from the connection
		log.Println("Server will try to get data from the teensy over TCP")
		go func() {
			for {
				dp, err := getDataPacket(conn)
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
		}()
	}

	// send commands to the teensy
	for {
		c := <-commandChan
		log.Println("Got command from the web client. Sending to teensy")
		_, err := conn.Write(c.WriteCommand())
		if err != nil {
			log.Print("Error sending command to teensy: ")
			log.Println(err)
			log.Println("Closing connection with teensy")
			conn.Close()
			return
		}
	}
}

func udpSocket() {
	udpConn, err := net.ListenUDP("udp", serverUDPAddress)
	if err != nil {
		panic(err)
	}

	go func() {
		<-abortChan
		for {
			udpConn.WriteToUDP([]byte("ABORT"), teensyUDPAddress)
		}
	}()

	for {
		dataPacket, err := getDataPacket(udpConn)
		if err != nil {

		}
		latestDataMutex.Lock()
		latestData = dataPacket
		latestDataMutex.Unlock()
	}
}

func getDataPacket(conn net.Conn) (*DataPacket, error) {
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
