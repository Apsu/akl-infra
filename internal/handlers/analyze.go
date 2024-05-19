package handlers

import (
	"net/http"

	"github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/api/pkg/analyzer/mini"
	"github.com/labstack/echo/v4"
)

func Analyze(ctx echo.Context) error {
	name := ctx.Param("name")
	layout, err := storage.Get(name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	mini.Table = mini.LoadTable()
	corpus := mini.LoadCorpus("corpora/monkeyracer/trigrams.json")
	counters := mini.Analyze(&layout, &corpus)

	return ctx.JSON(http.StatusOK, counters)
}
