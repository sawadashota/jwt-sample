package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sawadashota/jwt-sample/handler"
)

const DefaultPort = "8080"

var port string

func init() {
	port = os.Getenv("PORT")
	if port == "" {
		port = DefaultPort
	}
}

func main() {
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/auth", handler.RequireTokenAuthenticationHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
