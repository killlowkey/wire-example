package service

import (
	"context"
	"wire-example/internal/user/domain"
	"wire-example/internal/user/repository"
)

type Service struct {
	repo repository.UserRepository
}

func NewService(repo repository.UserRepository) Service {
	return Service{repo: repo}
}

func (s Service) FindUser(ctx context.Context, id int) (domain.User, error) {
	return s.repo.FindUser(ctx, id)
}
