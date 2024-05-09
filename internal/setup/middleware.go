package setup

import (
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Middleware(api *echo.Echo) {
	api.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
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

	api.Use(middleware.Recover())

	secureConfig := middleware.DefaultSecureConfig
	secureConfig.ContentSecurityPolicy = "default-src 'self'; script-src 'sha384-EfwldhYywH4qYH9vU8lMn+pd6pcH0kGpPUVJuwyHnj/5felkkIUVxf1wMAEX7rCY'; object-src 'none'; base-uri 'self'"
	secureConfig.ReferrerPolicy = "no-referrer"
	secureConfig.HSTSMaxAge = 31536000
	api.Use(middleware.SecureWithConfig(secureConfig))
}
