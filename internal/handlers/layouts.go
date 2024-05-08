package handlers

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func Layouts(ctx echo.Context) error {
	if entries, err := os.ReadDir("layouts"); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusNotFound)
	} else {
		var layouts []string
		for _, entry := range entries {
			if !entry.IsDir() {
				name := strings.TrimSuffix(entry.Name(), path.Ext(entry.Name()))
				layouts = append(layouts, name)
			}
		}
		return ctx.JSON(http.StatusOK, layouts)
	}
}
