package web_test

import (
	"testing"

	. "github.com/dnilosek/fib-overkill/worker/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig("redis://localhost", "8080", "in", "out")
	assert.Equal(t, "redis://localhost", cfg.DBURL)
	assert.Equal(t, "8080", cfg.DBPort)
	assert.Equal(t, "in", cfg.InputChannel)
	assert.Equal(t, "out", cfg.OutputChannel)
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()
	assert.Equal(t, "redis://localhost", cfg.DBURL)
	assert.Equal(t, "6379", cfg.DBPort)
	assert.Equal(t, "message", cfg.InputChannel)
	assert.Equal(t, "values", cfg.OutputChannel)
}
