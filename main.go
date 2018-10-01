package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sawadashota/jwt-sample/handler"
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
	http.HandleFunc("/login", handler.LoginHandler)
	http.HandleFunc("/auth", handler.RequireTokenAuthenticationHandler)

	fmt.Printf("Starting listen :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
