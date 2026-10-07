package application

import (
	"net/http"

	"github.com/Ficserbiyy/shortygo/internal/handler"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func (a *App) loadRoutes() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/health", func(c *echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	a.loadShortyRoutes(e.Group("/shorten"))

	a.router = e
}

func (a *App) loadShortyRoutes(e *echo.Group) {
	repo := handler.ShortyRepo{}

	e.POST("/", repo.Create())
}
