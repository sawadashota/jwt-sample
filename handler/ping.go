package handler

import "net/http"

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	w.Write([]byte("Pong"))
}
