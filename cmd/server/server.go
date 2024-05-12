package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/acme/autocert"
)

func Server(api *echo.Echo) {
	go func() {
		if val, ok := os.LookupEnv("AKL_PROD"); ok && val == "true" {
			api.AutoTLSManager.Cache = autocert.DirCache("/opt/cache")
			if err := api.StartAutoTLS(":443"); err != http.ErrServerClosed {
				log.Error(err)
			}
		} else {
			if err := api.Start(":8080"); err != http.ErrServerClosed {
				log.Error(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := api.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
