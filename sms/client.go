package sms

import (
	"gitea.com/yaothink/cloud/internal/kernel"
)

type Client struct {
	transport kernel.Transport
}

func NewClient(transport kernel.Transport) *Client {
	return &Client{
		transport: transport,
	}
}

func (c *Client) Delivery(request *DeliveryRequest) (response *DeliveryResponse, err error) {
	response = new(DeliveryResponse)

	return
}
