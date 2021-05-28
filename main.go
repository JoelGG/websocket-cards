package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joelgg/brag/router"
)

func main() {
	hub := router.NewHub()
	r := mux.NewRouter()
	r.HandleFunc("/connect/{roomid}", hub.NewConnection)
	r.HandleFunc("/checkroom/{roomid}", hub.LobbyExists)

	mux.CORSMethodMiddleware(r)

	http.ListenAndServe(":8080", r)
}
