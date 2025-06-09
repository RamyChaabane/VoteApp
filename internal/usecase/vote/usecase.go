package vote

import (
	"context"
)

type VoteRepository interface {
	IncrementVote(ctx context.Context, option string) error
}

type Service interface {
	Vote(ctx context.Context, option string) error
}

type voteService struct {
	repo VoteRepository
}

func NewService(repo VoteRepository) Service {
	return &voteService{repo: repo}
}

func (s *voteService) Vote(ctx context.Context, option string) error {
	return s.repo.IncrementVote(ctx, option)
}
