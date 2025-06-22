package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

//go:generate mockgen -source=./repo.go -destination=../../mocks/repo_mock.go -package=mocks

type VoteRepository interface {
	IncrementVote(ctx context.Context, option string) error
}

type VoteRepo struct {
	client *redis.Client
}

func NewVoteRepo() *VoteRepo {
	addr := os.Getenv("REDIS_HOST")
	if addr == "" {
		addr = "redis:6379"
	}

	client := redis.NewClient(&redis.Options{Addr: addr})
	return &VoteRepo{client: client}
}

func (r *VoteRepo) IncrementVote(ctx context.Context, option string) error {
	return r.client.Incr(ctx, option).Err()
}
