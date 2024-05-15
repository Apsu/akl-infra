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

func Server(srv *echo.Echo) {
	go func() {
		if val, ok := os.LookupEnv("AKL_PROD"); ok && val == "true" {
			srv.AutoTLSManager.Cache = autocert.DirCache("/opt/cache")
			srv.AutoTLSManager.Prompt = autocert.AcceptTOS
			srv.AutoTLSManager.HostPolicy = autocert.HostWhitelist("akl.gg")
			if err := srv.StartAutoTLS(":443"); err != http.ErrServerClosed {
				log.Error(err)
			}
		} else {
			if err := srv.Start(":8080"); err != http.ErrServerClosed {
				log.Error(err)
			}
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
