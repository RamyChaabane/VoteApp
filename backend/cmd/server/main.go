package main

import (
	"log"
	"net/http"

	"github.com/RamyChaabane/VoteApp/backend/internal/handler"
	"github.com/RamyChaabane/VoteApp/backend/internal/infrastructure/redis"
	usecase "github.com/RamyChaabane/VoteApp/backend/internal/usecase/vote"
)

func main() {
	// Dependencies
	repo := redis.NewVoteRepo()
	service := usecase.NewService(repo)
	voteHandler := handler.NewVoteHandler(service)

	// Routing
	http.Handle("/vote", voteHandler)

	log.Println("Backend listening on :8880")
	log.Fatal(http.ListenAndServe(":8880", nil))
}
