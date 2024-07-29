package app

import (
	"wire-example/service"
)

type App struct {
	Service *service.Service
}

func NewApp(service *service.Service) *App {
	return &App{Service: service}
}

func (a *App) Run() {
	a.Service.Serve()
}
