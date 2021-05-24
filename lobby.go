package main

import (
	"github.com/gorilla/websocket"
	"github.com/joelgg/brag/game"
)

type Lobby struct {
	clients   map[int]*Client
	msgs      chan WsMsg
	messages  []*Message
	index     int
	gameState game.GameState
}

func NewLobby() *Lobby {
	l := Lobby{}
	l.index = 0
	l.msgs = make(chan WsMsg)
	l.clients = map[int]*Client{}

	return &l
}

func (l *Lobby) Start() {
	for {
		m := <-l.msgs

		client := l.clients[m.idx]

		switch m.mt {
		case websocket.CloseMessage:
			client.Close()
		}

		message := NewMessage(string(m.msg), client)
		l.messages = append(l.messages, message)

		for x, cli := range l.clients {
			if x != m.idx {
				err := cli.conn.WriteMessage(m.mt, []byte(message.ToString()))
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func (l *Lobby) AddClient(c *websocket.Conn) {
	client := NewClient(l.index, c)
	l.clients[l.index] = client
	go client.Start(l.msgs)
	l.index++
}

func (l *Lobby) GetMessages() []string {
	mes := []string{}

	for _, m := range l.messages {
		mes = append(mes, m.ToString())
	}

	return mes
}
