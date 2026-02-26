package core

import (
	"gitea.com/yaothink/cloud/internal/internal/param"
	"gitea.com/yaothink/cloud/internal/kernel"
	"gitea.com/yaothink/cloud/sms"
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
