package main

import (
	"net/http"

	"github.com/akl-infra/akl.gg/internal/handlers"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Router() {
	api := echo.New()
	api.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Info("REQUEST", "uri", v.URI, "status", v.Status)
			} else {
				log.Error("REQUEST_ERROR", "uri", v.URI, "status", v.Status, "err", v.Error.Error())
			}
			return nil
		},
	}))

	// api.GET("/", handlers.Web)
	api.Static("/", "web")
	api.GET("/api", handlers.Api)
	api.GET("/api/layouts", handlers.Layouts)
	api.GET("/api/layout/:name", handlers.Layout)
	if err := api.Start(":80"); err != http.ErrServerClosed {
		log.Error(err)
	}
}

func main() {
	log.Info("Booting up")
	Router()
}
