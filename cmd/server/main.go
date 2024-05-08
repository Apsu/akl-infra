package main

import (
	"net/http"

	"github.com/apsu/akl-infra/internal/handlers"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func Router() {
	api := echo.New()
	api.GET("/", handlers.Root)
	api.GET("/layouts", handlers.Layouts)
	api.GET("/layout/:name", handlers.Layout)
	if err := api.Start(":42069"); err != http.ErrServerClosed {
		log.Error(err)
	}
}

func main() {
	log.Info("Booting up")
	Router()
}
