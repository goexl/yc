package sms

import (
	"github.com/goexl/yc/internal/constant"
)

type (
	MessageRequest struct {
		*request

		Template  uint64         `json:"templateId,string,omitempty"`
		Phone     string         `json:"phone,omitempty"`
		Arguments map[string]any `json:"arguments,omitempty"`
	}

	MessageResult struct {
		Phone   string `json:"phone,omitempty"`
		Success bool   `json:"success,omitempty"`
		Error   struct {
			Code    int    `json:"code,omitempty"`
			Message string `json:"message,omitempty"`
		} `json:"error,omitempty"`
	}

	MessageResponse struct {
		Result MessageResult `json:"result,omitempty"`
	}
)

func (*MessageRequest) Method() constant.Method {
	return constant.MethodPost
}

func (*MessageRequest) Url() string {
	return "messages"
}
