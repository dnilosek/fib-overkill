package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/dnilosek/fib-overkill/api/lib/web"
	"github.com/dnilosek/fib-overkill/lib/database"
)

func main() {
	app := web.Cli(&web.CliMethods{
		RunServer: runServer,
	})

	err := app.Run(os.Args)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func runServer(cfg *web.Config) error {

	redisConn := fmt.Sprintf("%s:%d", cfg.RedisURL, cfg.RedisPort)
	redis, err := database.OpenRedis(redisConn)
	if err != nil {
		return err
	}
	postgresConn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", cfg.PostgresURL, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)
	postgres, err := database.OpenPostgres(postgresConn)
	if err != nil {
		return err
	}
	server := web.NewServer(cfg, redis, postgres)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		handleInterrupt(server)
	}()

	err = server.Start()
	if err != nil {
		return err
	}

	wg.Wait()

	return nil
}

func handleInterrupt(server *web.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Interrupted...")
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	server.Stop(ctx)
}
