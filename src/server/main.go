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
	var testing *bool = flag.Bool("testing", false, "set to use testing mode")
	var debug *bool = flag.Bool("debug", false, "set to use debug mode. Debug mode lets you "+
		"communicate with the teensly solely over TCP for this weekends test event")
	var printData *bool = flag.Bool("log-data", false, "set to use log data as it comes in");
	flag.Parse()

	go startWebServer()
	if *debug {
		log.Println("Using debug mode. (For this weekend)")
		go startDebugArduinoComs()
	} else if *testing {
		log.Println("Using testing mode. (For vincent")
		go startFakeArduino()
	} else {
		go startArduinoComs()
	}
	if *printData {
		go startLogger()
	}
	for {}
}
