package ioc

import (
	"github.com/gin-gonic/gin"
	"wire-example/internal/order"
	"wire-example/internal/user"
	"wire-example/pkg/cron"
)

type App struct {
	userModule   *user.Module
	orderService *order.Service
	ginServer    *gin.Engine
	cron         *cron.Cron
}

func NewApp(
	userModule *user.Module,
	orderService *order.Service,
	ginServer *gin.Engine,
	cron *cron.Cron,
) *App {
	return &App{
		userModule:   userModule,
		orderService: orderService,
		ginServer:    ginServer,
		cron:         cron,
	}
}

func (a *App) Run() error {
	return a.ginServer.Run(":8080")
}
