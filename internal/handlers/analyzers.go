package handlers

import (
	"net/http"

	// "github.com/akl-infra/api/internal/storage"
	"github.com/akl-infra/api/pkg/analyzer/mini"
	"github.com/akl-infra/slf/v2"
	"github.com/labstack/echo/v4"
	"golang.org/x/exp/maps"
)

type AnalyzerFunc func(*slf.Layout, string) []float64

var AnalyzerFuncs = map[string]AnalyzerFunc{
	"mini": mini.Analyze,
}

func Analyzers(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, maps.Keys(AnalyzerFuncs))
}
