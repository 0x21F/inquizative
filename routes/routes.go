package routes

import (
	"context"
	"database/sql"

	"github.com/0x21F/inquizative/views/layouts"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteController struct {
	db *sql.DB
}

func (h *RouteController) ToServer() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Static("/public", "static")
	e.GET("/", h.index)
	// api := e.Group("/api/")

	return e
}

func (h *RouteController) index(c echo.Context) error {
	return layouts.Hero().Render(context.TODO(), c.Response().Writer)
}
