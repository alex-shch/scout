package main

import (
	log "alex-shch/scout/consolelog"

	"github.com/gorilla/websocket"
)

type Connection struct {
	ws *websocket.Conn
	// user data, player id?
}

func (self *Connection) recv() {
	for {
		_, msg, err := self.ws.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}

		log.Debugf("recv: %v", string(msg))

		// execute a command
		//self.Command(string(command))

		// update all conn
		//self.room.updateAll <- true
	}

	//self.room.leave <- pc
	self.ws.Close()
}

func (self *Connection) send(msg []byte) {
	err := self.ws.WriteMessage(websocket.TextMessage, msg)
	if err != nil {
		log.Error(err)
		//self.room.leave <- pc
		self.ws.Close()
	}
}

func newConnection(ws *websocket.Conn) Connection {
	c := Connection{ws}

	go c.recv()

	return c
}
