package handlers

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"

	"github.com/akl-infra/akl.gg/internal/types"
)

func Layout(ctx echo.Context) error {
	name := ctx.Param("name")
	if file, err := os.ReadFile("layouts/" + name + ".json"); err != nil {
		log.Error(err)
		return echo.NewHTTPError(http.StatusNotFound, err)
	} else {
		var layout types.Layout
		if err := json.Unmarshal(file, &layout); err != nil {
			log.Error(err)
			return echo.NewHTTPError(http.StatusInternalServerError, err)
		} else {
			return ctx.JSON(http.StatusOK, layout)
		}
	}
}
