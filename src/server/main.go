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
	flag.Parse()

	go startWebServer()
	if *testing {
		startFakeArduino()
	} else {
		startArduinoComs()
	}
}
