package param

import (
	"github.com/goexl/http"
	"github.com/goexl/log"
)

type Cloud struct {
	Id  string
	Key string

	Http   *http.Client
	Logger log.Logger
}

func NewCloud(id string, key string) *Cloud {
	return &Cloud{
		Id:  id,
		Key: key,

		Http:   http.New().Build(),
		Logger: log.New().Apply(),
	}
}
