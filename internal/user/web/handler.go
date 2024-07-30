package web

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"wire-example/internal/user/domain"
	"wire-example/internal/user/service"
)

type Handler struct {
	svc service.Service
}

func NewHandler(svc service.Service) *Handler {
	return &Handler{svc: svc}
}

func (h *Handler) PublicRoute(server *gin.Engine) {
	group := server.Group("/user")
	group.GET("/detail/:id", h.FindUserById)
}

func (h *Handler) FindUserById(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	user, _ := h.svc.FindUser(ctx.Request.Context(), id)
	ctx.JSONP(200, gin.H{"code": 200, "data": h.toUserVo(user)})
}

func (h *Handler) toUserVo(user domain.User) User {
	return User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
