package setup

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Middleware(srv *echo.Echo) {
	srv.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Info("REQUEST", "client", v.RemoteIP, "uri", v.URI, "status", v.Status)
			} else {
				log.Error("REQUEST_ERROR", "client", v.RemoteIP, "uri", v.URI, "status", v.Status, "err", v.Error.Error())
			}
			return nil
		},
	}))

	srv.Use(middleware.Recover())
	srv.Use(middleware.SecureWithConfig(middleware.SecureConfig{HSTSMaxAge: 3600}))
}
