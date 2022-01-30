package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (s Server) adminHomeView(c echo.Context) error {
	return c.Render(http.StatusOK, "base", nil)
}
