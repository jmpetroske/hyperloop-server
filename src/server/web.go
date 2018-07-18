package main

import (
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"strconv"
	"time"
)

const webServerAddress string = ":8080"

func allowCrossOrigin(w http.ResponseWriter) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Add("Access-Control-Allow-Methods", "POST")
	w.Header().Add("Access-Control-Expose-Headers", "Authorization")
	w.Header().Add("Access-Control-Max-Age", "600")
}

func missionHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

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

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))

	commandChan <- &MissionParamsCommand{
		Distance: float32(distance),
		Pressure: float32(pressure),
		TopSpeed: float32(topSpeed),
	}

	// log.Println("Got command: " + fmt.Sprintf("%+v" ,MissionParamsCommand{
	// 	Distance: float32(distance),
	// 	Pressure: float32(pressure),
	// 	TopSpeed: float32(topSpeed),
	// }))
}

func goToIdleHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

	commandChan <- &GoToIdleCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

func armHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

	commandChan <- &ArmCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

func startHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

	commandChan <- &StartCommand{}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

// see photonCommand.go for command values
func testCommandHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

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
	if command < 0 || command > 7 {
		http.Error(w, "400 - invalid command parameter, not in the valid range of values",
			http.StatusBadRequest)
		return
	}
	// log.Println("Got a command, sending to arduino coms")
	commandChan <- &TestingCommand{TestingCommandEnum(command)}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"success": true}`))
}

func abortHandler(w http.ResponseWriter, r *http.Request) {
	allowCrossOrigin(w)

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
		log.Println(err)
		return
	}

	for {
		var err error = nil
		latestDataMutex.Lock()
		if latestData != nil {
			err = webSocket.WriteJSON(*latestData)
		}
		latestDataMutex.Unlock()
		if err != nil {
			log.Println(err)
			return
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func startWebServer() {
	router := mux.NewRouter()
	router.HandleFunc("/mission", missionHandler).Methods("POST")
	router.HandleFunc("/goToIdle", goToIdleHandler).Methods("POST")
	router.HandleFunc("/arm", armHandler).Methods("POST")
	router.HandleFunc("/start", startHandler).Methods("POST")
	router.HandleFunc("/testCommand", testCommandHandler).Methods("POST")
	router.HandleFunc("/abort", abortHandler).Methods("POST")
	router.HandleFunc("/dataWebSocket", serveWebSocket)
	http.Handle("/", router)

	panic(http.ListenAndServe(webServerAddress, router))
}
