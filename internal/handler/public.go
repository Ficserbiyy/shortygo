package handler

import (
	"errors"
	"net/http"

	"github.com/Ficserbiyy/shortygo/internal/services/storage"
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

func (r *ShortyRepo) RedirectToURL() echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		shortcode := c.Param("shortcode")

		dbURL, err := storage.GetURLByShortcode(ctx, r.DB, shortcode)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return echo.NewHTTPError(http.StatusNotFound, "404 page not found")
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusFound, dbURL.URL)
	}
}
