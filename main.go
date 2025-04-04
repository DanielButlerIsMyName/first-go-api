package main

import (
	"log"
	"net/http"

	"first-go-api/handler"
)

func main() {
	handler.RegisterRoutes()

	log.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
