package main

import (
	// "fmt"
	"net/http"
	// "strings"
	"github.com/gorilla/mux"
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


func main() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler)
	router.HandleFunc("/arm", armHandler)
	router.HandleFunc("/start", startHandler)
	router.HandleFunc("/command", commandHandler)
	router.HandleFunc("/abort", abortHandler)
	http.Handle("/", router)

	panic(http.ListenAndServe(":8080", router))
}
