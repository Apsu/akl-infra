package main

import (
	"github.com/akl-infra/api/internal/auth"
	"github.com/akl-infra/api/internal/handlers"
	"github.com/akl-infra/api/internal/setup"
	"github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/api/pkg/analyzer/mini"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	log.Info("Booting up")

	site := echo.New()
	site.Pre(middleware.RemoveTrailingSlash())
	site.IPExtractor = echo.ExtractIPDirect()
	setup.Middleware(site)

	api := site.Host("api.akl.gg")

	// Public
	api.GET("/", handlers.Banner)
	api.GET("/layouts", handlers.Layouts)
	api.GET("/layout/:name", handlers.Layout)
	api.GET("/analyzers", handlers.Analyzers)
	api.GET("/analyze/:name", handlers.Analyze)

	// Protected
	api.PUT("/layout", handlers.AddLayout, auth.KeyAuth)

	storage.Init("layouts")
	// corpora.Init("corpora")
	mini.Init()
	Server(site)
}
