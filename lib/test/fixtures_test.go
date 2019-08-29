package test_test

import (
	"testing"

	. "github.com/dnilosek/kubing/apps/visitor-count/lib/test"
	"github.com/stretchr/testify/assert"
)

func TestMockDB(t *testing.T) {
	client := MockRedis()
	pong, err := client.Ping().Result()

	assert.Nil(t, err)
	assert.Equal(t, "PONG", pong)
}
