package builder

import (
	"github.com/goexl/log"
	"github.com/goexl/yc/internal/core"
	"github.com/goexl/yc/internal/internal/param"
)

type Cloud struct {
	params *param.Cloud
}

func NewCloud(id, key string) *Cloud {
	return &Cloud{
		params: param.NewCloud(id, key),
	}
}

func (c *Cloud) Logger(logger log.Logger) *Cloud {
	return c.set(func() {
		c.params.Logger = logger
	})
}

func (c *Cloud) Build() *core.Client {
	return core.NewClient(c.params)
}

func (c *Cloud) set(callback func()) (builder *Cloud) {
	callback()
	builder = c

	return
}
