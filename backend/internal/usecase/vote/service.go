package service

import (
	"context"

	"github.com/RamyChaabane/VoteApp/backend/internal/infrastructure/redis"
)

//go:generate mockgen -source=./service.go -destination=../../mocks/service_mock.go -package=mocks

type Service interface {
	Vote(ctx context.Context, option string) error
}

type voteService struct {
	repo redis.VoteRepository
}

func NewService(repo redis.VoteRepository) Service {
	return &voteService{repo: repo}
}

func (s *voteService) Vote(ctx context.Context, option string) error {
	return s.repo.IncrementVote(ctx, option)
}
