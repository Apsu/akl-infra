package handlers

import (
	"net/http"

	"github.com/akl-infra/akl.gg/internal/storage"
	"github.com/labstack/echo/v4"
)

func Layout(ctx echo.Context) error {
	name := ctx.Param("name")
	layout, err := storage.Get(name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	return ctx.JSON(http.StatusOK, layout)
}
