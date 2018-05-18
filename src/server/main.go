package main

import (
	"flag"
	"sync"
)

var latestDataMutex = &sync.Mutex{}
var latestData *DataPacket
var commandChan = make(chan PhotonCommand, 3)

func main() {
	var testing *bool = flag.Bool("testing", false, "set to use testing mode")
	var testing *bool = flag.Bool("debug", false, "set to use debug mode. Debug mode lets you " +
		"communicate with the teensly solely over TCP for this weekends test event")
	flag.Parse()

	go startWebServer()
	if *debug {
		startDebugArduino()
	} else if *testing {
		startFakeArduino()
	} else {
		startArduinoComs()
	}
}
