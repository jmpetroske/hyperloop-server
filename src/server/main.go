package main

import (
	"flag"
	"log"
	"sync"
)

var latestDataMutex = &sync.Mutex{}
var latestData *DataPacket
var commandChan = make(chan PhotonCommand, 20)
var abortChan = make(chan struct{}, 5)

func main() {
	var testing *bool = flag.Bool("testing", false, "set to use testing mode")
	var debug *bool = flag.Bool("debug", false, "set to use debug mode. Debug mode lets you "+
		"communicate with the teensly solely over TCP for this weekends test event")
	flag.Parse()

	go startWebServer()
	if *debug {
		log.Println("Using debug mode. (For this weekend)")
		startDebugArduinoComs()
	} else if *testing {
		log.Println("Using testing mode. (For vincent")
		startFakeArduino()
	} else {
		startArduinoComs()
	}
	for {}
}
