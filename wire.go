//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"wire-example/app"
	"wire-example/database"
	"wire-example/service"
)

func InitializeApp(dsn string) (*app.App, error) {
	wire.Build(
		database.NewDatabase,
		service.NewService,
		app.NewApp,
	)
	return &app.App{}, nil
}
