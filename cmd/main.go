package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func pingHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "pong\n")
}

func main() {
	http.HandleFunc("/ping", pingHandler)
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8090"
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
