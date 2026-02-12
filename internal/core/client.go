package core

import (
	"gitea.com/yaothink/cloud/internal/internal/param"
)

type Client struct {
	params *param.Cloud
}

func NewClient(params *param.Cloud) *Client {
	return &Client{
		params: params,
	}
}

func (c *Client) Sms() {

}
