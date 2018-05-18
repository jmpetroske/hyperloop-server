package main

import (
	"flag"
	"sync"
)

var latestDataMutex = &sync.Mutex{}
var latestData *DataPacket
var commandChan = make(chan PhotonCommand, 20)

func main() {
	var testing *bool = flag.Bool("testing", false, "set to use testing mode")
	var debug *bool = flag.Bool("debug", false, "set to use debug mode. Debug mode lets you "+
		"communicate with the teensly solely over TCP for this weekends test event")
	flag.Parse()

	go startWebServer()
	if *debug {
		startDebugArduinoComs()
	} else if *testing {
		startFakeArduino()
	} else {
		startArduinoComs()
	}
}
