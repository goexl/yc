package kernel

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goexl/exception"
	"github.com/goexl/gox"
	"github.com/goexl/gox/field"
	"github.com/goexl/yc/internal/internal/kernel"
	"github.com/goexl/yc/internal/internal/param"
)

type Transport struct {
	params *param.Cloud
}

func NewTransport(params *param.Cloud) *Transport {
	return &Transport{
		params: params,
	}
}

func (t *Transport) Do(ctx context.Context, req kernel.Request, rsp any) (err error) {
	request := t.params.Http.NewRequest()
	request.SetContext(ctx).SetResult(rsp)

	url := fmt.Sprintf("https://api.cloud.yaothink.tech%s", t.uri(req))
	fields := gox.Fields[any]{
		field.New("url", url),
	}

	t.params.Logger.Debug("开始调用接口", field.New("url", url), fields...)
	if bytes, me := json.Marshal(req); me != nil {
		err = me
	} else if pe := t.prepare(ctx, request, req, &bytes); nil != pe {
		err = pe
	} else if hpr, hpe := request.Post(url); nil != hpe {
		err = hpe
	} else if hpr.IsError() {
		err = t.handleException(req, hpr)
	}

	return
}

func (t *Transport) handleException(req kernel.Request, response *resty.Response) (err error) {
	if response.StatusCode() == http.StatusUnprocessableEntity {
		err = t.handleUnprocessableEntity(response)
	} else {
		err = t.handleOthers(req, response)
	}

	return
}

func (t *Transport) handleUnprocessableEntity(response *resty.Response) (err error) {
	exc := new(kernel.Exception)
	if ue := json.Unmarshal(response.Body(), exc); ue != nil {
		err = exception.New().Message("服务器数据返回格式错误").Field(field.New("body", string(response.Body()))).Build()
	} else if exc.Code == 1 {
		err = exception.New().Message("数组绑定出错").Field(field.New("code", 1)).Build()
	} else if exc.Code == 2 {
		err = exception.New().Message("设置数据默认值出错").Field(field.New("code", 2)).Build()
	} else if exc.Code == 3 {
		err = exception.New().Message("数组校验不通过").Field(field.New("code", 3)).Build()
	} else {
		err = exception.New().Message(exc.Message).Field(field.New("code", exc.Code), field.New("data", exc.Data)).Build()
	}

	return
}

func (t *Transport) handleOthers(req kernel.Request, response *resty.Response) (err error) {
	exc := new(kernel.Exception)
	if ue := json.Unmarshal(response.Body(), exc); ue != nil {
		err = exception.New().Message("服务器数据返回格式错误").Field(field.New("body", string(response.Body()))).Build()
	} else if exc.Code == 1 {
		err = exception.New().Message("数组绑定出错").Field(field.New("code", 1)).Build()
	} else if exc.Code == 2 {
		err = exception.New().Message("设置数据默认值出错").Field(field.New("code", 2)).Build()
	} else if exc.Code == 3 {
		err = exception.New().Message("数组校验不通过").Field(field.New("code", 3)).Build()
	} else {
		err = exception.New().Message("未知错误").Field(field.New("body", string(response.Body()))).Build()
	}

	return
}

func (t *Transport) prepare(
	_ context.Context,
	request *resty.Request,
	req kernel.Request, payload *[]byte,
) (err error) {
	id := t.params.Id
	key := t.params.Key
	host := "api.cloud.yaothink.tech"
	algorithm := "YC-ZONGLIANG"
	contentType := "application/json; charset=utf-8"
	category := req.Category()
	product := req.Product()
	function := req.Function()
	timestamp := time.Now().Unix()

	// 步骤一：构建规范请求
	method := req.Method()
	uri := t.uri(req)
	query := ""
	headers := fmt.Sprintf("content-type:%s\nhost:%s\n", contentType, host)
	signedHeaders := "content-type;host"
	hashedRequestPayload := t.sha256Hex(*payload)
	canonicalRequest := fmt.Sprintf(
		"%s\n%s\n%s\n%s\n%s\n%s",
		method,
		uri,
		query,
		headers,
		signedHeaders,
		hashedRequestPayload,
	)
	t.params.Logger.Debug("构建规范请求完成", field.New("result", canonicalRequest))

	// 步骤二：构建待签名字符串
	date := time.Unix(timestamp, 0).UTC().Format("2006-01-02")
	scope := fmt.Sprintf("%s/%s/%s/%s/yc_request", date, category, product, function)
	hashedCanonicalRequest := t.sha256Hex([]byte(canonicalRequest))
	stringToSign := fmt.Sprintf(
		"%s\n%d\n%s\n%s\n%s\n%s\n%s",
		algorithm,
		timestamp,
		category,
		product,
		function,
		scope,
		hashedCanonicalRequest,
	)
	t.params.Logger.Debug("构建待签名字符串完成", field.New("result", stringToSign))

	// 步骤三：计算签名
	secretDate := t.hmacSHA256(date, algorithm+key)
	secretService := t.hmacSHA256(fmt.Sprintf("%s/%s/%s", category, product, function), secretDate)
	secretSigning := t.hmacSHA256("yc_request", secretService)
	signature := hex.EncodeToString([]byte(t.hmacSHA256(stringToSign, secretSigning)))
	t.params.Logger.Debug("计算签名完成", field.New("result", signature))

	token := fmt.Sprintf(
		"%s,%s,%s,%s,%s,%s",
		id,
		category,
		product,
		function,
		signedHeaders,
		signature,
	)
	t.params.Logger.Debug("填充授权标头", field.New("token", token))
	request.SetAuthScheme(algorithm).SetAuthToken(token)
	request.SetBody(*payload)
	request.SetHeader("X-Yc-Timestamp", strconv.FormatInt(timestamp, 10))
	request.SetHeader("Content-Type", contentType)

	return
}

func (t *Transport) uri(req kernel.Request) string {
	return fmt.Sprintf(
		"/categories/%s/products/%s/functions/%s/%s",
		req.Category(), req.Product(), req.Function(), req.Url(),
	)
}

func (t *Transport) sha256Hex(from []byte) (to string) {
	sum := sha256.Sum256(from)
	to = hex.EncodeToString(sum[:])

	return
}

func (t *Transport) hmacSHA256(from, key string) (to string) {
	hashed := hmac.New(sha256.New, []byte(key))
	hashed.Write([]byte(from))
	to = string(hashed.Sum(nil))

	return
}
