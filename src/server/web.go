package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
)

const webServerAddress string = ":8080"

func missionHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("distance") == "" ||
		r.FormValue("pressure") == "" ||
		r.FormValue("topSpeed") == "" {
		http.Error(w, "400 - missing a mission parameter", http.StatusBadRequest)
		return
	}

	distance, err := strconv.ParseFloat(r.FormValue("distance"), 32)
	if err != nil {
		log.Println(err)
	}
	pressure, err := strconv.ParseFloat(r.FormValue("pressure"), 32)
	if err != nil {
		log.Println(err)
	}
	topSpeed, err := strconv.ParseFloat(r.FormValue("topSpeed"), 32)
	if err != nil {
		log.Println(err)
	}
	commandChan <- &MissionParamsCommand{
		Distance: float32(distance),
		Pressure: float32(pressure),
		TopSpeed: float32(topSpeed),
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success: true}"))
}

func armHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &ArmCommand{}
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &StartCommand{}
}

func commandHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &TestingCommand{ /* TODO */ }
}

func abortHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &AbortCommand{}
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

	for {
		if err = webSocket.WriteJSON(DataPacket{}); err != nil {
			log.Println(err)
		}
	}
}

func startWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler).Methods("POST")
	router.HandleFunc("/arm", armHandler).Methods("POST")
	router.HandleFunc("/start", startHandler).Methods("POST")
	router.HandleFunc("/command", commandHandler).Methods("POST")
	router.HandleFunc("/abort", abortHandler).Methods("POST")
	router.HandleFunc("/dataWebSocket", serveWebSocket)
	http.Handle("/", router)

	panic(http.ListenAndServe(webServerAddress, router))
}
