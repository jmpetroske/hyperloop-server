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
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success: true}"))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &StartCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success: true}"))
}

/*
 * valid command values:
 * 0: engageBreaks
 * 1: disengageBreaks
 * 2: engageSolenoids
 * 3: disengageSolenoids
 * 4: engageBallValves
 * 5: disengageBallValves
 */
func commandHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("command") == "" {
		http.Error(w, "400 - missing the command parameter", http.StatusBadRequest)
		return
	}
	command, err := strconv.ParseInt(r.FormValue("command"), 10, 32)
	if err != nil {
		log.Println(err)
		http.Error(w, "400 - invalid command parameter. Pass an int", http.StatusBadRequest)
		return
	}
	if command < 0 || command > 5 {
		http.Error(w, "400 - invalid command parameter, not in the valid range of values",
			http.StatusBadRequest)
		return
	}
	commandChan <- &TestingCommand{TestingCommandEnum(command)}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success: true}"))
}

func abortHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &AbortCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{success: true}"))
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
