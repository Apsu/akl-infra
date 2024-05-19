package handlers

import (
	"net/http"

	"github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/api/pkg/analyzer"
	"github.com/labstack/echo/v4"
)

func Analyzers(ctx echo.Context) error {

	return ctx.JSON(http.StatusOK)
}
