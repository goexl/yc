package kernel

import (
	"github.com/goexl/yc/internal/constant"
)

type Request interface {
	Category() string

	Product() string

	Function() string

	Url() string

	Method() constant.Method
}
