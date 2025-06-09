package handler

import (
	"fmt"
	"github.com/RamyChaabane/VoteApp/backend/internal/infrastructure"
	"net/http"

	usecase "github.com/RamyChaabane/VoteApp/backend/internal/usecase/vote"
)

var voteUseCase = usecase.UseCase{
	Repo: infrastructure.NewVoteRepo(),
}

func VoteHandler(w http.ResponseWriter, r *http.Request) {
	option := r.FormValue("vote")
	if option == "" {
		http.Error(w, "Missing vote", http.StatusBadRequest)
		return
	}

	err := voteUseCase.Vote(r.Context(), option)
	if err != nil {
		http.Error(w, "Vote failed", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Vote for %s recorded!", option)
}
