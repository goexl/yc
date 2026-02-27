package core

import (
	"github.com/goexl/yc/internal/internal/param"
	"github.com/goexl/yc/internal/kernel"
	"github.com/goexl/yc/sms"
)

type Client struct {
	params *param.Cloud
}

func NewClient(params *param.Cloud) *Client {
	return &Client{
		params: params,
	}
}

func (c *Client) Sms() *sms.Client {
	return sms.NewClient(kernel.NewTransport(c.params))
}
