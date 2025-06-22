package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RamyChaabane/VoteApp/backend/internal/domain/vote"
	usecase "github.com/RamyChaabane/VoteApp/backend/internal/usecase/vote"
)

//go:generate mockgen -source=./vote_handler.go -destination=../mocks/vote_handler_mock.go -package=mocks

type VoteInterface interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type VoteHandler struct {
	service usecase.Service
}

func NewVoteHandler(service usecase.Service) VoteInterface {
	return &VoteHandler{service: service}
}

func (h *VoteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("starting vote handler")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	option := r.FormValue("vote")

	log.Printf("vote option: %s", option)

	if !vote.IsValidOption(option) {
		errMsg := "Invalid vote option"
		w.Header().Set("Content-Length", fmt.Sprint(len(errMsg)))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errMsg))
		return
	}

	err := h.service.Vote(r.Context(), option)
	if err != nil {
		errMsg := "Vote failed"
		w.Header().Set("Content-Length", fmt.Sprint(len(errMsg)))
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errMsg))
		return
	}

	msg := fmt.Sprintf("Vote for %s recorded!", option)
	w.Header().Set("Content-Length", fmt.Sprint(len(msg)))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(msg))
	return
}
