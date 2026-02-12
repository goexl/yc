package sms

type (
	DeliveryRequest struct {
		*request

		Template uint64         `json:"templateId,string,omitempty"`
		Phones   []string       `json:"phones,omitempty"`
		Values   map[string]any `json:"varName2Values,omitempty"`
	}

	DeliveryResponse struct {
	}
)
