package ioc

import (
	"wire-example/pkg/db"
)

func NewDatabase() (*db.Database, error) {
	return &db.Database{DSN: "tcp://username:password@127.0.0.1:3306"}, nil
}
