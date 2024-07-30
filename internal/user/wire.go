//go:build wireinject

package user

import (
	"github.com/google/wire"
	"sync"
	"wire-example/internal/user/repository"
	"wire-example/internal/user/service"
	"wire-example/internal/user/web"
	"wire-example/pkg/cache"
	"wire-example/pkg/db"
)

type (
	Handler = web.Handler
	Service = service.Service
)

// InitModule 要是 db 和 cache 这两个依赖不在 user 包下，则需要从外面引入，需要放入到这个方法参数下
func InitModule(db *db.Database, cache cache.ICache) *Module {
	wire.Build(InitService, InitHandler, wire.Struct(new(Module), "*"))
	return new(Module)
}

func InitUserRepo(db *db.Database) repository.UserRepository {
	return repository.NewUserRepository(db)
}

var (
	once = &sync.Once{}
	svc  service.Service
)

func InitService(db *db.Database) service.Service {
	once.Do(func() {
		orderRepository := repository.NewUserRepository(db)
		svc = service.NewService(orderRepository)
	})
	return svc
}

func InitHandler(service service.Service) *Handler {
	wire.Build(web.NewHandler)
	return new(Handler)
}
