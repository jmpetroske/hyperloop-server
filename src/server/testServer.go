package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const webServerAddress string = ":8080"

var commandChan = make(chan PhotonCommand, 3)
var testData DataPacket = DataPacket{
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
	w.Write([]byte(`{"success": true}`))
}

func armHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &ArmCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &StartCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
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
	w.Write([]byte(`{"success": true}`))
}

func abortHandler(w http.ResponseWriter, r *http.Request) {
	commandChan <- &AbortCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
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
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler).Methods("POST")
	router.HandleFunc("/arm", armHandler).Methods("POST")
	router.HandleFunc("/start", startHandler).Methods("POST")
	router.HandleFunc("/command", commandHandler).Methods("POST")
	router.HandleFunc("/abort", abortHandler).Methods("POST")
	router.HandleFunc("/dataWebSocket", serveWebSocket)
	http.Handle("/", router)

	go func() {
		for {
			log.Println(<-commandChan)
		}
	}()
	panic(http.ListenAndServe(webServerAddress, router))
}
