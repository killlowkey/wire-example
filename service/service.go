package service

import (
	"fmt"
	"wire-example/database"
)

type Service struct {
	DB *database.Database
}

func NewService(db *database.Database) *Service {
	return &Service{DB: db}
}

func (s *Service) Serve() {
	fmt.Println("Service is serving with database:", s.DB.DSN)
}
