//go:build wireinject
// +build wireinject

package ioc

import (
	"github.com/google/wire"
	"wire-example/internal/order"
	"wire-example/internal/user"
	"wire-example/pkg/cron"
)

var BaseSet = wire.NewSet(NewDatabase, NewCache)

func InitializeApp() (*App, error) {
	wire.Build(
		wire.Struct(new(App), "*"),
		BaseSet,
		order.NewOrderService,            // 构造器注入
		wire.Struct(new(cron.Cron), "*"), // 往该 struct 注入所有字段
		user.InitModule,
		wire.FieldsOf(new(*user.Module), "Handler"), // GinServer 依赖这个 handler，需要从 Module 提取出来，注入到 Gin Server
		NewGinServer,                                //需要放到最后面
	)
	return new(App), nil
}
