// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/app"
	"wire-example/cache"
	"wire-example/database"
	"wire-example/service"
)

// Injectors from wire.go:

func InitializeApp(dsn string) (*app.App, error) {
	databaseDatabase, err := database.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	iCache := cache.NewRedisCache()
	userService := service.NewUserService(databaseDatabase, iCache)
	orderService := service.NewOrderService(databaseDatabase, iCache)
	appApp := app.NewApp(userService, orderService)
	return appApp, nil
}

// wire.go:

var BaseSet = wire.NewSet(database.NewDatabase, cache.NewRedisCache)
