package handler

import (
	"errors"
	"net/http"

	"github.com/Ficserbiyy/shortygo/internal/model"
	"github.com/Ficserbiyy/shortygo/internal/services/storage"
	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
)

type ShortyRepo struct {
	DB *pgx.Conn
}

var (
	errNotFound = echo.NewHTTPError(http.StatusNotFound, "404 URL not found")
)

// Create method creates a new short URL.
func (s *ShortyRepo) Create() echo.HandlerFunc {
	return func(c *echo.Context) error {
		r := c.Request()

		// Parse form data
		err := r.ParseForm()
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		url := model.URL{
			URL: r.FormValue("link"),
		}
		ctx := r.Context()

		shortcode, err := storage.Create(ctx, s.DB, url)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusCreated, map[string]string{"shortcode": shortcode})
	}
}

// RedirectToURL method redirects to original URL.
func (s *ShortyRepo) RedirectToURL() echo.HandlerFunc {
	return func(c *echo.Context) error {
		ctx := c.Request().Context()
		shortcode := c.Param("shortcode")

		dbURL, err := storage.GetByShortcode(ctx, s.DB, shortcode)
		if err != nil {
			if errors.Is(err, pgx.ErrNoRows) {
				return errNotFound
			}
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.Redirect(http.StatusFound, dbURL.URL)
	}
}
