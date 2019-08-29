package web

import (
	"log"
	"strings"

	"github.com/urfave/cli"
)

const (
	portArg             = "port"
	apiPathArg          = "api-path"
	redisUrlArg         = "redis-url"
	redisPortArg        = "redis-port"
	postgresUrlArg      = "postgres-url"
	postgresPortArg     = "postgres-port"
	postgresDatabaseArg = "postgres-database"
	postgresUserArg     = "postgres-user"
	postgresPasswordArg = "postgres-password"
)

// Define operations that CLI impliments
type RunServer func(cfg *Config) error
type CliMethods struct {
	RunServer RunServer
}

// Create the CLI app
func Cli(methods *CliMethods) *cli.App {
	app := cli.NewApp()

	// Define the app parameters and flags
	app.Name = "fib-api"
	app.HelpName = "fib-api"
	app.Usage = "Run the fibonacci calculator api"
	app.UsageText = "fib-api -p API_PORT --api-path API_PATH --redis-url REDIS_HOST --redis-port REDIS_PORT --postgres-url PGHOST --postgres-port PGPORT --postgres-database PGDATABASE --postgres-user PGUSER --postgres-password PGPASSWORD"
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:   strings.Join([]string{portArg, "p"}, ","),
			Value:  defaultAPIPort,
			Usage:  "port to listen on",
			EnvVar: "API_PORT",
		},
		cli.StringFlag{
			Name:   apiPathArg,
			Value:  defaultAPIPath,
			Usage:  "url path prefix for mounting API router",
			EnvVar: "API_PATH",
		},
		cli.StringFlag{
			Name:   redisUrlArg,
			Value:  defaultRedisURL,
			Usage:  "Connection URL to redis server",
			EnvVar: "REDIS_HOST",
		},
		cli.IntFlag{
			Name:   redisPortArg,
			Value:  defaultRedisPort,
			Usage:  "Redis server port",
			EnvVar: "REDIS_PORT",
		},
		cli.StringFlag{
			Name:   postgresUrlArg,
			Value:  defaultPostgresURL,
			Usage:  "Connection URL to postgres server",
			EnvVar: "PGHOST",
		},
		cli.IntFlag{
			Name:   postgresPortArg,
			Value:  defaultPostgresPort,
			Usage:  "Postgres server port",
			EnvVar: "PGPORT",
		},
		cli.StringFlag{
			Name:   postgresDatabaseArg,
			Value:  defaultPostgresDB,
			Usage:  "Postgres server port",
			EnvVar: "PGDATABASE",
		},
		cli.StringFlag{
			Name:   postgresUserArg,
			Value:  defaultPostgresUser,
			Usage:  "Postgres user",
			EnvVar: "PGUSER",
		},
		cli.StringFlag{
			Name:   postgresPasswordArg,
			Value:  defaultPostgresPassword,
			Usage:  "Postgres user password",
			EnvVar: "PGPASSWORD",
		},
	}

	// Create the action for the app
	app.Action = func(c *cli.Context) error {

		cfg := getConfig(c)

		log.Printf("API PATH: 		%s", cfg.APIPath)
		log.Printf("REDIS URL:		%s", cfg.RedisURL)
		log.Printf("POSTGRES URL:	%s", cfg.PostgresURL)

		return methods.RunServer(cfg)
	}
	return app
}

func getConfig(c *cli.Context) *Config {
	port := c.Int(portArg)
	apiPath := c.String(apiPathArg)
	redisURL := c.String(redisUrlArg)
	redisPort := c.Int(redisPortArg)
	postgresURL := c.String(postgresUrlArg)
	postgresPort := c.Int(postgresPortArg)
	postgresDB := c.String(postgresDatabaseArg)
	postgresUser := c.String(postgresUserArg)
	postgresPassword := c.String(postgresPasswordArg)

	return NewConfig(port, apiPath, redisURL, redisPort, postgresURL, postgresPort, postgresDB, postgresUser, postgresPassword)
}
