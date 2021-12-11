package server

import "github.com/gorilla/mux"

func (w *WS) setRoutes(r *mux.Router) {
	r.HandleFunc("/ws", w.connHandler)
}
