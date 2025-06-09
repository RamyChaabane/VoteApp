package main

import (
	"log"
	"net/http"

	"github.com/RamyChaabane/VoteApp/backend/internal/handler"
)

func main() {
	http.HandleFunc("/vote", handler.VoteHandler)
	log.Println("Backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
