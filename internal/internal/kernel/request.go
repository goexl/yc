package kernel

import (
	"gitea.com/yaothink/cloud/internal/constant"
)

type Request interface {
	Category() string

	Product() string

	Function() string

	Url() string

	Method() constant.Method
}
