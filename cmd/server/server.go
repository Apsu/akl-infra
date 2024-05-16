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

func Server(site *echo.Echo) {
	go func() {
		if val, ok := os.LookupEnv("AKL_PROD"); ok && val == "true" {
			site.AutoTLSManager.Cache = autocert.DirCache("/opt/cache")
			site.AutoTLSManager.Prompt = autocert.AcceptTOS
			site.AutoTLSManager.HostPolicy = autocert.HostWhitelist("akl.gg", "api.akl.gg")
			if err := site.StartAutoTLS(":443"); err != http.ErrServerClosed {
				log.Error(err)
			}
		} else {
			if err := site.Start(":8080"); err != http.ErrServerClosed {
				log.Error(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := site.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
