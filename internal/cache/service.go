package cache

import (
	"context"
	"errors"

	"github.com/redis/go-redis/v9"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Set(ctx context.Context, key string, value string, ttl int) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if value == "" {
		return errors.New("value cannot be empty")
	}
	return s.repo.Set(ctx, key, value, ttl)
}

func (s *Service) Get(ctx context.Context, key string) (string, error) {
	if key == "" {
		return "", errors.New("key cannot be empty")
	}
	value, err := s.repo.Get(ctx, key)

	if err == redis.Nil {
		return "", errors.New("key not found")
	}
	if err != nil {
		return "", err
	}
	return value, nil
}

func (s *Service) Delete(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	return s.repo.Delete(ctx, key)
}
