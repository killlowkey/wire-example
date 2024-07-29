package database

import "fmt"

type Database struct {
	DSN string
}

func NewDatabase(dsn string) (*Database, error) {
	// 模拟数据库连接
	fmt.Println("Connecting to database with DSN:", dsn)
	return &Database{DSN: dsn}, nil
}
