package application

import "github.com/labstack/echo/v5"

type App struct {
	router *echo.Echo
}

func NewApp() *App {
	app := &App{}

	app.loadRoutes()
	return app
}

func (a *App) Start() {

	if err := a.router.Start(":8080"); err != nil {
		a.router.Logger.Error("failed to start server", "error", err)
	}
}
