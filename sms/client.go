package sms

import (
	"context"

	"github.com/goexl/yc/internal/kernel"
)

type Client struct {
	transport *kernel.Transport
}

func NewClient(transport *kernel.Transport) *Client {
	return &Client{
		transport: transport,
	}
}

func (c *Client) Message(ctx context.Context, request *MessageRequest) (response *MessageResponse, err error) {
	response = new(MessageResponse)
	err = c.transport.Do(ctx, request, response)

	return
}

func (c *Client) Delivery(ctx context.Context, request *DeliveryRequest) (response *DeliveryResponse, err error) {
	response = new(DeliveryResponse)
	err = c.transport.Do(ctx, request, response)

	return
}
