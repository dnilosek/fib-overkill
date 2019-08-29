package web

// Defaults
var (
	defaultAPIPort          = 8080
	defaultAPIPath          = "/"
	defaultRedisURL         = "redis://localhost"
	defaultRedisPort        = 6379
	defaultPostgresURL      = "postgresql://localhost"
	defaultPostgresPort     = 5432
	defaultPostgresDB       = "postgres"
	defaultPostgresUser     = "postgres"
	defaultPostgresPassword = "postgres_password"
)

// Config type
type Config struct {
	APIPort          int
	APIPath          string
	RedisURL         string
	RedisPort        int
	PostgresURL      string
	PostgresPort     int
	PostgresDB       string
	PostgresUser     string
	PostgresPassword string
}

// Create config
func NewConfig(apiPort int, apiPath, redisURL string, redisPort int, postgresURL string, postgresPort int, postgresDB, postgresUser, postgresPassword string) *Config {
	return &Config{
		APIPort:          apiPort,
		APIPath:          apiPath,
		RedisURL:         redisURL,
		RedisPort:        redisPort,
		PostgresURL:      postgresURL,
		PostgresPort:     postgresPort,
		PostgresDB:       postgresDB,
		PostgresUser:     postgresUser,
		PostgresPassword: postgresPassword,
	}
}

// Create default
func DefaultConfig() *Config {
	return NewConfig(
		defaultAPIPort,
		defaultAPIPath,
		defaultRedisURL,
		defaultRedisPort,
		defaultPostgresURL,
		defaultPostgresPort,
		defaultPostgresDB,
		defaultPostgresUser,
		defaultPostgresPassword)
}
