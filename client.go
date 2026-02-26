package cloud

import (
	"gitea.com/yaothink/cloud/internal/builder"
	"gitea.com/yaothink/cloud/internal/core"
)

// Client 客户端
type Client = core.Client

func New(id string, key string) *builder.Cloud {
	return builder.NewCloud(id, key)
}
