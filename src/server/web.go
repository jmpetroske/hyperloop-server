package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"net/http"
)

const webServerAddress string = ":8080"

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
}

func serveWebSocket(w http.ResponseWriter, r *http.Request) {
	webSocket, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	webSocket.WriteMessage(websocket.TextMessage, []byte("Hello Vincent"))
}

func startWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler)
	router.HandleFunc("/arm", armHandler)
	router.HandleFunc("/start", startHandler)
	router.HandleFunc("/command", commandHandler)
	router.HandleFunc("/abort", abortHandler)
	router.HandleFunc("/dataWebSocket", abortHandler)
	http.Handle("/", router)

	panic(http.ListenAndServe(webServerAddress, router))
}
