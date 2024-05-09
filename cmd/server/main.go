package main

import (
	"net/http"

	// "github.com/akl-infra/akl.gg/internal/auth"
	"github.com/akl-infra/akl.gg/internal/handlers"
	"github.com/akl-infra/akl.gg/internal/setup"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/acme/autocert"
)

func main() {
	log.Info("Booting up")

	api := echo.New()
	setup.Middleware(api)

	api.AutoTLSManager.Cache = autocert.DirCache("/opt/cache")

	api.Static("/", "web")

	// protected := api.Group("/api")
	// protected.Use(middleware.KeyAuth(auth.TokenValidator))

	api.GET("/api", handlers.Banner)
	api.GET("/api/layouts", handlers.Layouts)
	api.GET("/api/layout/:name", handlers.Layout)

	if err := api.StartAutoTLS(":443"); err != http.ErrServerClosed {
		log.Error(err)
	}
}
