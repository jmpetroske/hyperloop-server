package main

import (
	"fmt"
	"net/http"
	// "strings"
	"github.com/gorilla/mux"
	"net"
	// "github.com/gorilla/websocket"
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
		fmt.Println("%b", conn)
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler)
	router.HandleFunc("/arm", armHandler)
	router.HandleFunc("/start", startHandler)
	router.HandleFunc("/command", commandHandler)
	router.HandleFunc("/abort", abortHandler)
	http.Handle("/", router)

	go tcpSocket();
	
	panic(http.ListenAndServe(":8080", router))
}
