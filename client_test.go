package yc_test

import (
	"testing"

	"github.com/goexl/yc"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, yc.New("id", "key").Build())
}
