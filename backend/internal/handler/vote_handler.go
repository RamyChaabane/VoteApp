package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/RamyChaabane/VoteApp/backend/internal/domain/vote"
	usecase "github.com/RamyChaabane/VoteApp/backend/internal/usecase/vote"
)

type VoteHandler struct {
	Service usecase.Service
}

func NewVoteHandler(service usecase.Service) *VoteHandler {
	return &VoteHandler{Service: service}
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
		http.Error(w, "Invalid vote option", http.StatusBadRequest)
		return
	}

	err := h.Service.Vote(r.Context(), option)
	if err != nil {
		http.Error(w, "Vote failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Vote for %s recorded!", option)
}
