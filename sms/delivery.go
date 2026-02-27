package sms

import (
	"github.com/goexl/yc/internal/constant"
)

type (
	DeliveryRequest struct {
		*request

		Template uint64         `json:"templateId,string,omitempty"`
		Phones   []string       `json:"phones,omitempty"`
		Values   map[string]any `json:"varName2Values,omitempty"`
	}

	DeliverResult struct {
		Phone   string `json:"phone,omitempty"`
		Success bool   `json:"success,omitempty"`
		Error   string `json:"error,omitempty"`
	}

	DeliveryResponse struct {
		Results []DeliverResult `json:"results,omitempty"`
	}
)

func (*DeliveryRequest) Method() constant.Method {
	return constant.MethodPost
}

func (*DeliveryRequest) Url() string {
	return "api/sms/deliveries"
}
