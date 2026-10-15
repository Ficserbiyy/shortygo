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

	a.loadShortyRoutes(e)

	a.router = e
}

func (a *App) loadShortyRoutes(e *echo.Echo) {
	repo := handler.ShortyRepo{
		DB: a.db,
	}
	g := e.Group("/shorten")

	e.GET("/url/:shortcode", repo.RedirectToURL())

	g.Any("/:shortcode", repo.LinkPage())
	g.POST("/", repo.Create())
}
