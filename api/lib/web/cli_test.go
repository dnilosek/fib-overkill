package web_test

import (
	"testing"

	. "github.com/dnilosek/fib-overkill/api/lib/web"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCli(t *testing.T) {
	// Mock server
	runServer := func(*Config) error { return nil }

	cli := Cli(&CliMethods{
		RunServer: runServer,
	})

	assert.Equal(t, "fib-api", cli.Name)
	assert.Equal(t, "fib-api", cli.HelpName)
	assert.Equal(t, "Run the fibonacci calculator api", cli.Usage)
	assert.Equal(t, "fib-api -p API_PORT --api-path API_PATH --redis-url REDIS_HOST --redis-port REDIS_PORT --postgres-url PGHOST --postgres-port PGPORT --postgres-database PGDATABASE --postgres-user PGUSER --postgres-password PGPASSWORD", cli.UsageText)
	assert.Equal(t, 0, len(cli.Commands))
	assert.Equal(t, 9, len(cli.Flags))
	assert.Equal(t, "port,p", cli.Flags[0].GetName())
	assert.Equal(t, "api-path", cli.Flags[1].GetName())
	assert.Equal(t, "redis-url", cli.Flags[2].GetName())
	assert.Equal(t, "redis-port", cli.Flags[3].GetName())
	assert.Equal(t, "postgres-url", cli.Flags[4].GetName())
	assert.Equal(t, "postgres-port", cli.Flags[5].GetName())
	assert.Equal(t, "postgres-database", cli.Flags[6].GetName())
	assert.Equal(t, "postgres-user", cli.Flags[7].GetName())
	assert.Equal(t, "postgres-password", cli.Flags[8].GetName())
}

func TestCliAction(t *testing.T) {

	var result *Config
	runServer := func(cfg *Config) error { result = cfg; return nil }

	cli := Cli(&CliMethods{
		RunServer: runServer,
	})

	// Test defaults
	err := cli.Run([]string{"app"})
	require.Nil(t, err)
	assert.Equal(t, 8080, result.APIPort)
	assert.Equal(t, "/", result.APIPath)
	assert.Equal(t, "redis://localhost", result.RedisURL)
	assert.Equal(t, 6379, result.RedisPort)
	assert.Equal(t, "postgresql://localhost", result.PostgresURL)
	assert.Equal(t, 5432, result.PostgresPort)
	assert.Equal(t, "postgres", result.PostgresDB)
	assert.Equal(t, "postgres", result.PostgresUser)
	assert.Equal(t, "postgres_password", result.PostgresPassword)

	// Test input
	err = cli.Run([]string{"app", "--port=80", "--api-path=/api", "--redis-url=redis://localhost", "--redis-port=8080", "--postgres-url=postgres://localhost", "--postgres-port=8080", "--postgres-database=test", "--postgres-user=me", "--postgres-password=bad"})
	require.Nil(t, err)
	assert.Equal(t, 80, result.APIPort)
	assert.Equal(t, "/api", result.APIPath)
	assert.Equal(t, "redis://localhost", result.RedisURL)
	assert.Equal(t, 8080, result.RedisPort)
	assert.Equal(t, "postgres://localhost", result.PostgresURL)
	assert.Equal(t, 8080, result.PostgresPort)
	assert.Equal(t, "test", result.PostgresDB)
	assert.Equal(t, "me", result.PostgresUser)
	assert.Equal(t, "bad", result.PostgresPassword)
}
