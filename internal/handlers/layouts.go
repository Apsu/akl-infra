package handlers

import (
	"net/http"

	"github.com/akl-infra/akl.gg/internal/storage"
	"github.com/labstack/echo/v4"
)

func Layouts(ctx echo.Context) error {
	layouts := storage.List()
	return ctx.JSON(http.StatusOK, layouts)
}
