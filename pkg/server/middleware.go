package server

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// requestID sets the request ID to the echo context value "requestID", if a request ID header is not present this
// middleware generates one
func (s Server) requestID(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		reqIDHeader := c.Request().Header.Get("X-Request-ID")
		if reqIDHeader == "" {
			c.Set("requestID", uuid.New().String())
		} else {
			c.Set("requestID", reqIDHeader)
		}
		return next(c)
	}
}
