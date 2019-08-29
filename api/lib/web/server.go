package web

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dnilosek/fib-overkill/lib/database"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type Server struct {
	*Config
	router     *echo.Echo
	redisDB    *database.RedisDB
	postgresDB *database.PostgresDB
}

func NewServer(cfg *Config, redisDB *database.RedisDB, postgresDB *database.PostgresDB) *Server {

	// Create new server
	server := &Server{
		Config:     cfg,
		router:     echo.New(),
		redisDB:    redisDB,
		postgresDB: postgresDB,
	}

	// Create endpoints
	server.router.GET("/", server.Index)
	server.router.GET("/values/all", server.GetAllValues)
	server.router.GET("/values/current", server.GetCurrentValues)
	server.router.POST("/values", server.PostValue)

	// Clear startup info printing
	server.router.HideBanner = true
	server.router.HidePort = true

	// Attach render function
	server.router.Use(middleware.Logger())
	server.router.Use(middleware.CORS())

	return server
}

// Start serving
func (server *Server) Start() error {
	server.postgresDB.Client.Exec("CREATE TABLE IF NOT EXISTS values (number INT)")
	return server.router.Start(fmt.Sprintf(":%d", server.APIPort))
}

// Stop serving
func (server *Server) Stop(ctx context.Context) error {
	return server.router.Shutdown(ctx)
}

// Serve the index page
func (server *Server) Index(context echo.Context) error {
	return context.String(http.StatusOK, "OK")
}

// Serve the index page
func (server *Server) GetAllValues(context echo.Context) error {
	rows, err := server.postgresDB.Client.Query("SELECT * from values")
	if err != nil {
		return err
	}
	vals := make([]int, 0)
	for rows.Next() {
		var val int
		err := rows.Scan(&val)
		if err != nil {
			return err
		}
		vals = append(vals, val)
	}
	return context.JSON(http.StatusOK, vals)
}

// Serve the index page
func (server *Server) GetCurrentValues(context echo.Context) error {
	vals, err := server.redisDB.HGetAll("values")
	if err != nil {
		return err
	}
	return context.JSON(http.StatusOK, vals)
}

func (server *Server) PostValue(context echo.Context) error {
	idx := context.FormValue("index")
	if v, err := strconv.Atoi(idx); v > 100 || err != nil {
		return err
	}
	err := server.redisDB.HSet("values", idx, "placeholder")
	if err != nil {
		return err
	}
	server.redisDB.Client.Publish("message", idx)
	_, err = server.postgresDB.Client.Exec("INSERT INTO values(number) VALUES($1)", idx)
	if err != nil {
		return err
	}

	return nil
}
