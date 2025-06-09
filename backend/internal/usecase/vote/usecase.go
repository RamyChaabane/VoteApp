package vote

import (
	"context"
)

type VoteRepository interface {
	IncrementVote(ctx context.Context, option string) error
}

type UseCase struct {
	Repo VoteRepository
}

func (uc *UseCase) Vote(ctx context.Context, option string) error {
	return uc.Repo.IncrementVote(ctx, option)
}
