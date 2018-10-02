package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/sawadashota/jwt-sample/middleware"

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
	auth := middleware.New(middleware.Auth)

	mux := http.NewServeMux()

	mux.HandleFunc("/login", post(handler.LoginHandler))
	mux.HandleFunc("/hello", auth.Then(get(handler.PingHandler)))

	fmt.Printf("Starting listen :%s...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), mux))
}

func post(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Bad Request Method", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}

func get(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Bad Request Method", http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
