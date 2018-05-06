package main

var currDataPacket *DataPacket

func logDataPacket(dp *DataPacket) {
	currDataPacket = dp
	// TODO actually log it somewhere and keep old data
}

func getMostRecentData(retDataPacket *DataPacket) {
	*retDataPacket = *currDataPacket
}
