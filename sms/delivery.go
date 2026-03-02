package sms

import (
	"github.com/goexl/exception"
	"github.com/goexl/gox/http"
	"github.com/goexl/yc/internal/constant"
)

type (
	DeliveryRequest struct {
		*request

		Template  uint64           `json:"templateId,string,omitempty"`
		Phones    []string         `json:"phones,omitempty"`
		Arguments []map[string]any `json:"arguments,omitempty"`
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
	return "deliveries"
}
