package app

import (
	"wire-example/service"
)

type App struct {
	userService  *service.UserService
	orderService *service.OrderService
}

func NewApp(userService *service.UserService, orderService *service.OrderService) *App {
	return &App{
		userService:  userService,
		orderService: orderService,
	}
}

func (a *App) Run() {
	a.userService.Serve()
	a.orderService.Serve()
}
