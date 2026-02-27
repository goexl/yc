package yc

import (
	"github.com/goexl/yc/internal/builder"
	"github.com/goexl/yc/internal/core"
)

// Client 客户端
type Client = core.Client

func New(id string, key string) *builder.Cloud {
	return builder.NewCloud(id, key)
}
