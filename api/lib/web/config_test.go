package web_test

import (
	"testing"

	. "github.com/dnilosek/fib-overkill/api/lib/web"
	"github.com/stretchr/testify/assert"
)

func TestNewConfig(t *testing.T) {
	cfg := NewConfig(80, "/api", "redis://localhost", 8080, "postgres://localhost", 8081, "test", "me", "bad")

	assert.Equal(t, 80, cfg.APIPort)
	assert.Equal(t, "/api", cfg.APIPath)
	assert.Equal(t, "redis://localhost", cfg.RedisURL)
	assert.Equal(t, 8080, cfg.RedisPort)
	assert.Equal(t, "postgres://localhost", cfg.PostgresURL)
	assert.Equal(t, 8081, cfg.PostgresPort)
	assert.Equal(t, "test", cfg.PostgresDB)
	assert.Equal(t, "me", cfg.PostgresUser)
	assert.Equal(t, "bad", cfg.PostgresPassword)
}

func TestDefaultConfig(t *testing.T) {
	cfg := DefaultConfig()

	assert.Equal(t, 8080, cfg.APIPort)
	assert.Equal(t, "/", cfg.APIPath)
	assert.Equal(t, "redis://localhost", cfg.RedisURL)
	assert.Equal(t, 6379, cfg.RedisPort)
	assert.Equal(t, "postgresql://localhost", cfg.PostgresURL)
	assert.Equal(t, 5432, cfg.PostgresPort)
	assert.Equal(t, "postgres", cfg.PostgresDB)
	assert.Equal(t, "postgres", cfg.PostgresUser)
	assert.Equal(t, "postgres_password", cfg.PostgresPassword)

}
