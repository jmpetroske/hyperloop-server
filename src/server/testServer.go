package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"math/rand"
	"net/http"
	"time"
)

const webServerAddress string = ":8080"

var testData DataPacket

type Employee struct {
	firstName string
}

func missionHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func armHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

func abortHandler(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	webSocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		testData.Timestamp += 100000
		testData.Speed += (rand.Float32() - 0.5) * 0.3
		testData.Distance += testData.Speed * 0.1

		if err = webSocket.WriteJSON(testData); err != nil {
			fmt.Println(err)
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	testData = DataPacket{
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
	
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler)
	router.HandleFunc("/arm", armHandler)
	router.HandleFunc("/start", startHandler)
	router.HandleFunc("/command", commandHandler)
	router.HandleFunc("/abort", abortHandler)
	router.HandleFunc("/dataWebSocket", serveWebSocket)
	http.Handle("/", router)

	panic(http.ListenAndServe(webServerAddress, router))
}
