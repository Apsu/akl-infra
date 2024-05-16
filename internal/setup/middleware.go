package setup

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Middleware(site *echo.Echo) {
	site.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:   true,
		LogURI:      true,
		LogError:    true,
		HandleError: true, // forwards error to the global error handler, so it can decide appropriate status code
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			if v.Error == nil {
				log.Info("REQUEST", "client", c.RealIP(), "uri", v.URI, "status", v.Status)
			} else {
				log.Error("REQUEST_ERROR", "client", c.RealIP(), "uri", v.URI, "status", v.Status, "err", v.Error.Error())
			}
			return nil
		},
	}))

	site.Use(middleware.Recover())
	site.Use(middleware.SecureWithConfig(middleware.SecureConfig{HSTSMaxAge: 3600}))
}
