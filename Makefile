run:
	go run src/server/ardiuno.go src/server/dataLog.go src/server/dataPacket.go src/server/fakeArduino.go src/server/main.go src/server/photonCommand.go src/server/web.go  -tcp-only -log-data

