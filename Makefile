main:
	go run src/server/ardiuno.go src/server/dataLog.go src/server/dataPacket.go src/server/main.go src/server/photonCommand.go src/server/web.go

testing:
	go run src/server/testServer.go src/server/dataPacket.go
