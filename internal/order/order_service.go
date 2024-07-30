package order

import (
	"context"
	"fmt"
	"wire-example/pkg/cache"
	"wire-example/pkg/db"
)

type Service struct {
	db    *db.Database
	cache cache.ICache
}

func NewOrderService(db *db.Database, cache cache.ICache) *Service {
	return &Service{
		db:    db,
		cache: cache,
	}
}

func (o *Service) Serve() {
	fmt.Println("OrderService is serving with database:", o.db.DSN)
	val, _ := o.cache.Get(context.TODO(), "OrderService")
	fmt.Println("OrderService is serving with cache:", val)
}
