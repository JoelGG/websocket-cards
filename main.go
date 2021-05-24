package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	hub := NewHub()
	r := mux.NewRouter()
	r.HandleFunc("/connect/{roomid}", hub.NewConnection)
	r.HandleFunc("/checkroom/{roomid}", hub.LobbyExists)

	mux.CORSMethodMiddleware(r)

	http.ListenAndServe(":8080", r)
}
