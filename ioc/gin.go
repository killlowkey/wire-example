package ioc

import (
	"github.com/gin-gonic/gin"
	"wire-example/internal/user/web"
)

func NewGinServer(userHandler *web.Handler) *gin.Engine {
	engine := gin.Default()
	userHandler.PublicRoute(engine)
	return engine
}
