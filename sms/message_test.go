package sms_test

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/goexl/gox/rand"
	"github.com/goexl/yc"
	"github.com/goexl/yc/sms"
)

func TestMessage(t *testing.T) {
	id := os.Getenv("SECRET_ID")
	key := os.Getenv("SECRET_KEY")
	phone := os.Getenv("PHONE")

	client := yc.New(id, key).Build().Sms()
	if rsp, err := client.Message(context.Background(), &sms.MessageRequest{
		Template: 268942876933623808,
		Phone:    phone,
		Arguments: map[string]any{
			"code": rand.New().String().Code().Build().Generate(),
		},
	}); err != nil {
		t.Error(err)
	} else if !rsp.Success {
		t.Error(rsp.Error)
	} else {
		fmt.Println(rsp)
	}
}
