package main

import (
	"fmt"
	"time"

	log "alex-shch/scout/consolelog"

	"github.com/googollee/go-socket.io"
)

const _GAME_ROOM = "game"

type Game struct {
	server  *socketio.Server
	players map[socketio.Socket]Player
}

func NewGame(server *socketio.Server) *Game {
	game := &Game{server: server, players: map[socketio.Socket]Player{}}

	server.On("connection", func(so socketio.Socket) {
		log.Info("on connection")
		game.AddPlayer(so)
	})

	server.On("error", func(so socketio.Socket, err error) {
		log.Error("error:", err)
	})

	return game
}

func (self *Game) AddPlayer(so socketio.Socket) {
	player := Player{Id: fmt.Sprintf("%p", so)}
	self.players[so] = player

	so.Join(_GAME_ROOM)

	so.On("move", func(msg string) {
		//log.Info("emit:", so.Emit("chat message", msg))
		//so.BroadcastTo("chat", "chat message", msg)
		log.Infof("move cmd: %v", msg)
	})

	so.On("disconnection", func() {
		log.Info("on disconnect")

		// TODO сделать безопасно (параллельный доступ!!!)
		player := self.players[so]
		delete(self.players, so)
		so.BroadcastTo(_GAME_ROOM, "playerRemoved", player.Id)
	})

	so.BroadcastTo(_GAME_ROOM, "playerAdded", player.Id)
}

func (self *Game) Loop() {
	for {
		msg := fmt.Sprintf("%v", time.Now())

		//log.Debugf("tick %v", msg)
		self.server.BroadcastTo(_GAME_ROOM, "tick", msg)

		time.Sleep(300 * time.Millisecond)
	}
}
