package main

func main() {
	var dataPacketChan chan DataPacket = make(chan DataPacket, 10)
	go startArduinoComs(dataPacketChan)
	startWebServer(dataPacketChan)
}
