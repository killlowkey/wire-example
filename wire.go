//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/app"
	"wire-example/cache"
	"wire-example/cron"
	"wire-example/database"
	"wire-example/service"
)

var BaseSet = wire.NewSet(database.NewDatabase, cache.NewRedisCache)

func InitializeApp(dsn string) (*app.App, error) {
	wire.Build(
		BaseSet,
		wire.Struct(new(cron.Cron), "*"), // 往该 struct 注入所有字段
		service.NewUserService,           // 构造器注入
		service.NewOrderService,
		app.NewApp,
	)
	return &app.App{}, nil
}
