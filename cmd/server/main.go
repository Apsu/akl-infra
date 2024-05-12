package main

import (
	// "github.com/akl-infra/akl.gg/internal/auth"
	"github.com/akl-infra/akl.gg/internal/handlers"
	"github.com/akl-infra/akl.gg/internal/setup"
	"github.com/akl-infra/akl.gg/internal/storage"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Info("Booting up")

	api := echo.New()
	setup.Middleware(api)

	// protected := api.Group("/api")
	// protected.Use(middleware.KeyAuth(auth.TokenValidator))

	api.GET("/api", handlers.Banner)
	api.GET("/api/layouts", handlers.Layouts)
	api.GET("/api/layout/:name", handlers.Layout)

	storage.Init("layouts")
	Server(api)
}
