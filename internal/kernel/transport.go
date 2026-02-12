package kernel

import (
	"context"
	"fmt"

	"gitea.com/yaothink/cloud/internal/internal/kernel"
	"gitea.com/yaothink/cloud/internal/internal/param"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
)

type Transport struct {
	params *param.Cloud
}

func NewTransport(params *param.Cloud) *Transport {
	return &Transport{
		params: params,
	}
}

func (t *Transport) Do(ctx context.Context, req kernel.Request, rsp any) (code uint32, err error) {
	response := new(dto.Response)
	response.Result = rsp

	request := t.http.NewRequest()
	request.SetContext(ctx).SetBody(req).SetResult(response)

	url := fmt.Sprintf(
		"https://aip.cloud.yaotink.tech/categories/%s/services/%s/functions/%s",
		req.Category(), req.Service(), req.Function(),
	)
	fields := gox.Fields[any]{
		field.New("url", url),
	}
	if token, pte := t.pickToken(ctx); nil != pte {
		err = pte
	} else if hpr, hpe := request.SetQueryParam("access_token", token).Post(url); nil != hpe {
		err = hpe
	} else if hpr.IsError() {
		bodyField := field.New("body", string(hpr.Body()))
		message := "百度服务器返回错误"
		err = exception.New().Code(1).Message(message).Field(bodyField, fields...).Build()
		t.logger.Error(message, bodyField, fields...)
	} else if response.IsError() {
		bodyField := field.New("body", string(hpr.Body()))
		message := "接口调用出错"
		t.logger.Warn(message, bodyField, fields...)
	}

	// 将代码回传给上级调用
	code = response.Code

	return
}
