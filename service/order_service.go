package service

import (
	"context"
	"fmt"
	"wire-example/cache"
	"wire-example/database"
)

type OrderService struct {
	db    *database.Database
	cache cache.ICache
}

func NewOrderService(db *database.Database, cache cache.ICache) *OrderService {
	return &OrderService{
		db:    db,
		cache: cache,
	}
}

func (o *OrderService) Serve() {
	fmt.Println("OrderService is serving with database:", o.db.DSN)
	val, _ := o.cache.Get(context.TODO(), "OrderService")
	fmt.Println("OrderService is serving with cache:", val)
}
