// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"wire-example/app"
	"wire-example/database"
	"wire-example/service"
)

// Injectors from wire.go:

func InitializeApp(dsn string) (*app.App, error) {
	databaseDatabase, err := database.NewDatabase(dsn)
	if err != nil {
		return nil, err
	}
	serviceService := service.NewService(databaseDatabase)
	appApp := app.NewApp(serviceService)
	return appApp, nil
}