package main

import (
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/labstack/echo/v4"
)

type LayoutKey struct {
	X      uint `json:"x"`
	Y      uint `json:"y"`
	Row    uint `json:"row"`
	Col    uint `json:"col"`
	Finger uint `json:"finger"`
}

type LayoutVariant struct {
	Uid  int64                `json:"uid"`
	Link string               `json:"link"`
	Keys map[string]LayoutKey `json:"keys"`
}

type Keys map[string]LayoutKey
type Layout map[string]LayoutVariant
type Layouts map[string]Layout

var layouts = Layouts{
	"qwerty": {
		"standard": {
			Uid: 123456789,
			Keys: Keys{
				"q": {
					Row:    3,
					Col:    0,
					Finger: 0,
				},
			},
		},
	},
}

func rRoot(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "AKL AKL AKL!")
}

func rLayouts(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, layouts)
}

func rLayout(ctx echo.Context) error {
	name := ctx.Param("name")
	variant := ctx.QueryParam("variant")
	if variant == "" {
		variant = "standard"
	}
	if layout, ok := layouts[name]; ok {
		if data, ok := layout[variant]; ok {
			return ctx.JSON(http.StatusOK, data)
		} else {
			return echo.NewHTTPError(http.StatusNotFound, "Variant not found")
		}
	} else {
		return echo.NewHTTPError(http.StatusNotFound, "Layout not found")
	}
}

func Router() {
	api := echo.New()
	api.GET("/", rRoot)
	api.GET("/layouts", rLayouts)
	api.GET("/layout/:name", rLayout)
	if err := api.Start(":42069"); err != http.ErrServerClosed {
		log.Error(err)
	}
}

func main() {
	log.Info("Booting up")
	Router()
}
