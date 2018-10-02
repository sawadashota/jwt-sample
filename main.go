package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sawadashota/jwt-sample/handler"
	"github.com/sawadashota/jwt-sample/middleware"
	"github.com/sawadashota/jwt-sample/route"
)

const (
	DefaultPort = "8080"
)

var (
	port string
)

func init() {
	port = getEnvString("APP_PORT", DefaultPort)
}

func getEnvString(key, defaultValue string) string {
	v := os.Getenv(key)
	if v == "" {
		v = defaultValue
	}

	return v
}

func main() {
	mux := http.NewServeMux()

	middleware.New(mux).Group(
		route.Post("/login", handler.LoginHandler),
	)

	middleware.New(mux, middleware.Auth).Group(
		route.Get("/hello", handler.PingHandler),
		route.Get("/ping", handler.PingHandler),
	)

	log.Printf("Starting listen :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}
