package main

import (
	log "alex-shch/scout/consolelog"

	"net/http"

	"github.com/gorilla/websocket"
)

var connections []Connection = []Connection{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		return
	}

	conn := newConnection(ws)
	connections = append(connections, conn)
}

func main() {
	log.Info("start")
	defer log.Info("stop")

	log.SetLogLevel(log.DEBUG)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Debug(r.URL.Path)
		http.ServeFile(w, r, "www"+r.URL.Path)
	})

	http.HandleFunc("/ws", wsHandler)

	go GameLoop()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Error("ListenAndServe: ", err)
		panic(err)
	}
}
