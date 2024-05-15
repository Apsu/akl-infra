package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func AddLayout(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}
