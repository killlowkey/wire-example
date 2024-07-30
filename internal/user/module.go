package user

import (
	"wire-example/pkg/cache"
)

type Module struct {
	Handler *Handler // 引用的是 wire_gen.go
	Svc     Service  // 引用的是 wire_gen.go
	Cache   cache.ICache
}
