package main

import (
	"fmt"
	"time"
)

const logRate = 1

func startLogger() {
	for {
		latestDataMutex.Lock()
		fmt.Println(*latestData)
		latestDataMutex.Unlock()
		time.Sleep(logRate * time.Second)
	}
}
