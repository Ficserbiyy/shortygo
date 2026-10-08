package handler

import (
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
)

type ShortyRepo struct {
	DB *pgx.Conn
}

// Create method creates a new short URL.
func (r *ShortyRepo) Create() echo.HandlerFunc {
	return func(c *echo.Context) error {
		return c.NoContent(http.StatusNotImplemented)
	}
}
