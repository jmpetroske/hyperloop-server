package main

import (
	"log"
	"math/rand"
	"time"
)

var testData DataPacket = DataPacket{
	0,
	3,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0.0,
	0,
	0,
	0,
	false,
	false,
	false,
	false,
	false,
}

func startFakeArduino() {
	go func() {
		for {
			c := <-commandChan
			testData.Mode = c.ExpectedNextMode()
			log.Println(c)
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
