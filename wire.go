//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/app"
	"wire-example/cache"
	"wire-example/database"
	"wire-example/service"
)

var BaseSet = wire.NewSet(database.NewDatabase, cache.NewRedisCache)

func InitializeApp(dsn string) (*app.App, error) {
	wire.Build(
		BaseSet,
		service.NewUserService,
		service.NewOrderService,
		app.NewApp,
	)
	return &app.App{}, nil
}
