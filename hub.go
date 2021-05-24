package main

import (
	"bytes"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type Hub struct {
	lobbys   map[string]*Lobby
	index    int
	upgrader websocket.Upgrader
}

func NewHub() *Hub {
	h := Hub{}
	h.lobbys = map[string]*Lobby{}
	h.index = 0
	h.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
	return &h
}

func (h *Hub) LobbyExists(w http.ResponseWriter, r *http.Request) {
	roomvar := mux.Vars(r)["roomid"]
	contains := h.lobbys[roomvar] != nil
	buf := bytes.NewBuffer([]byte{})
	buf.WriteString(strconv.FormatBool(contains))
	buf.WriteTo(w)
}

func (h *Hub) NewConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := h.upgrader.Upgrade(w, r, nil)
	log.Println("New connection")
	if err != nil {
		log.Panicln(err)
	}

	roomvar := mux.Vars(r)["roomid"]

	lobby := h.lobbys[roomvar]

	if lobby != nil {
		lobby.AddClient(conn)
		conn.WriteMessage(websocket.TextMessage, []byte("Joined room "+roomvar))
		conn.WriteMessage(websocket.TextMessage, []byte("Previous messages: \n"+strings.Join(lobby.GetMessages(), "\n")))
	} else {
		lob := h.addLobby(roomvar)
		lob.AddClient(conn)
		go lob.Start()
		conn.WriteMessage(websocket.TextMessage, []byte("Room created "+roomvar))
	}
}

func (h *Hub) addLobby(name string) *Lobby {
	lob := NewLobby()
	h.lobbys[name] = lob
	h.index++
	return lob
}

func (h *Hub) disconnectClient(index int, conn *websocket.Conn, name string) {
	log.Print("Client disconnected")
	conn.Close()
	delete(h.lobbys[name].clients, index)
	if len(h.lobbys[name].clients) == 0 {
		delete(h.lobbys, name)
	}
}
