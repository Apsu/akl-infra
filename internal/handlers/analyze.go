package handlers

import (
	// "fmt"
	"net/http"
	// "strings"

	"github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/api/pkg/analyzer/mini"
	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

func Analyze(ctx echo.Context) error {
	name := ctx.Param("name")
	corpus := ctx.Param("corpus")
	if corpus == "" {
		corpus = "monkeyracer"
	}
	log.Info(corpus)
	layout, err := storage.Get(name)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, err)
	}

	// mini.Table = mini.LoadTable()
	// corpus := mini.LoadCorpus("corpora/monkeyracer/trigrams.json")
	counters := mini.Analyze(&layout, corpus)

	// var sb strings.Builder
	metrics := make(map[string]float64)

	for i, counter := range counters {

		metrics[mini.MetricToStr(mini.Metric(i))] = counter
		// fmt.Fprintf(&sb, "%s: %f\n", mini.MetricToStr(mini.Metric(i)), counter)
	}

	// return ctx.JSON(http.StatusOK, sb.String())
	return ctx.JSON(http.StatusOK, counters)
}
