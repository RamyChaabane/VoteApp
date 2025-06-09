package handler

import (
	"fmt"
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
	option := r.FormValue("vote")
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
