package repository

import (
	"context"
	"fmt"
	"wire-example/internal/user/domain"
	"wire-example/pkg/db"
)

type UserRepository interface {
	FindUser(ctx context.Context, id int) (domain.User, error)
}

type UserRepositoryImpl struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) UserRepository {
	return &UserRepositoryImpl{db: db}
}

func (u *UserRepositoryImpl) FindUser(ctx context.Context, id int) (domain.User, error) {
	return domain.User{ID: id, Name: fmt.Sprintf("user-%d", id)}, nil
}
