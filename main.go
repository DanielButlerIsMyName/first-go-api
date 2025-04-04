package main

import (
	"log"
	"net/http"
	"first-go-api/handler"
)

func main() {
	http.HandleFunc("/calc", handler.Calculate)

	log.Println("Server starting on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
