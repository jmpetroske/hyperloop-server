package main

import (
	"fmt"
	"net/http"
	// "strings"
	"net"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

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

func tcpSocket() {
	readBuf := make([]byte, 2048)
	
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		// handle error
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		fmt.Println("Successfully connected");
		fmt.Println("%b", conn);
		bytesRead, err := conn.Read(readBuf);
		fmt.Println(bytesRead)
	}
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

	webSocket.WriteMessage(websocket.TextMessage, []byte("Hello"));
}

func startWebClient() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler)
	router.HandleFunc("/arm", armHandler)
	router.HandleFunc("/start", startHandler)
	router.HandleFunc("/command", commandHandler)
	router.HandleFunc("/abort", abortHandler)
	router.HandleFunc("/dataWebSocket", abortHandler)
	http.Handle("/", router)

	panic(http.ListenAndServe(":8080", router))
}

func main() {
	// TODO init
	go tcpSocket();
	startWebClient();
}
