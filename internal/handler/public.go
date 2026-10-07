package handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

type ShortyRepo struct{}

// Create method creates a new short URL.
func (r *ShortyRepo) Create() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.NoContent(http.StatusNotImplemented)
	}
}
