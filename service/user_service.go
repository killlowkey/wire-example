package service

import (
	"context"
	"fmt"
	"wire-example/cache"
	"wire-example/database"
)

type UserService struct {
	db    *database.Database
	cache cache.ICache
}

func NewUserService(db *database.Database, cache cache.ICache) *UserService {
	return &UserService{
		db:    db,
		cache: cache,
	}
}

func (s *UserService) Serve() {
	fmt.Println("userService is serving with database:", s.db.DSN)
	val, _ := s.cache.Get(context.TODO(), "UserService")
	fmt.Println("UserService is serving with cache:", val)
}
