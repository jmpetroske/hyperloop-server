package main

import (
	"math/rand"
	"log"
	"time"
)

var testData DataPacket = DataPacket{
	0,
	3,
	88.82,
	65.83,
	56.13,
	76.29,
	53.54,
	63.85,
	95.88,
	13.46,
	51.1,
	15.8,
	10.73,
	59.6,
	65.0,
	78.79,
	63.15,
	68.73,
	96.57,
	true,
	true,
	true,
	false,
	4.38,
	27.92,
	20.33,
	7.21,
	48.20,
	88.49,
	40.0,
	75.71,
	78.43,
	78.41,
	28.81,
	58.13,
	90.10,
	87.90,
	76.49,
	34.64,
	12.4,
	96.29,
	1,
	0,
	1,
	false}

func startFakeArduino() {
	go func() {
		for {
			log.Println(<-commandChan)
		}
	}()
	
	latestDataMutex.Lock()
	latestData = &testData

	latestDataMutex.Unlock()
	
	for {
		latestDataMutex.Lock()
		testData.Timestamp += 100000
		testData.Speed += (rand.Float32() - 0.2) * 0.3
		testData.Distance += testData.Speed * 0.1
		latestDataMutex.Unlock()
		time.Sleep(100 * time.Millisecond)
	}
}
