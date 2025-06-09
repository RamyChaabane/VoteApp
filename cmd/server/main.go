package main

import (
	"log"
	"net/http"

	"github.com/RamyChaabane/VoteApp/internal/handler"
	"github.com/RamyChaabane/VoteApp/internal/infrastructure/redis"
	usecase "github.com/RamyChaabane/VoteApp/internal/usecase/vote"
)

func main() {
	// Dependencies
	repo := redis.NewVoteRepo()
	service := usecase.NewService(repo)
	voteHandler := handler.NewVoteHandler(service)

	// Routing
	http.Handle("/vote", voteHandler)

	log.Println("Backend listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
