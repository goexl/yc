package cloud_test

import (
	"testing"

	"gitea.com/yaothink/cloud"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	assert.NotNil(t, cloud.New("id", "key").Build())
}
