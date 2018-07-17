package main

import (
	"flag"
	"log"
	"sync"
)

var latestDataMutex = &sync.Mutex{}
var latestData *DataPacket = &DataPacket{}
var commandChan = make(chan PhotonCommand, 20)
var abortChan = make(chan struct{}, 5)

func main() {
	var tcpOnly *bool = flag.Bool("tcp-only", false, "Use tcp for all communications")
	var logData *bool = flag.Bool("log-data", false,
		"Log data as it comes in from the teensy. Data will be printed to stdout")
	var vincentTest *bool = flag.Bool("vincent-testing", false, "vincent's mode")
	flag.Parse()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		startWebServer()
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		if *vincentTest {
			log.Println("Using test mode for vincent")
			startFakeArduino()
		} else if *tcpOnly {
			log.Println("Using only tcp for communications")
			tcpComs(true)
		} else {
			go udpSocket()
			log.Println("Using TCP and UDP for communications")
			tcpComs(false)
		}
		wg.Done()
	}()

	if *logData {
		go startLogger()
	}

	wg.Wait()
}
