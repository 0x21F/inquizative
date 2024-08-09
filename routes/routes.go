package routes

import (
	"database/sql"

	"github.com/labstack/echo/v4"
)

type RouteController struct {
	db *sql.DB
}

func NewController() *RouteController {
	res := &RouteController{}

	return res
}

func (h *RouteController) Index(c echo.Context) error {

	return nil
}
