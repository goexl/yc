package kernel

type Exception struct {
	Code    uint8  `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
