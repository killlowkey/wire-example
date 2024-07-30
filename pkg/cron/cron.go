package cron

import (
	"context"
	"fmt"
	"wire-example/pkg/cache"
)

type Cron struct {
	Cache cache.ICache // 字段注入必须要 export 的
}

func (c *Cron) Create(ctx context.Context, name string) error {
	fmt.Println("create cron: ", name)
	return nil
}
