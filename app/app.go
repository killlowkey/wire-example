package app

import (
	"context"
	"wire-example/cron"
	"wire-example/service"
)

type App struct {
	userService  *service.UserService
	orderService *service.OrderService
	cron         *cron.Cron
}

func NewApp(userService *service.UserService, orderService *service.OrderService, cron *cron.Cron) *App {
	return &App{
		userService:  userService,
		orderService: orderService,
		cron:         cron,
	}
}

func (a *App) Run() {
	a.userService.Serve()
	a.orderService.Serve()
	a.cron.Create(context.TODO(), "task")
}
