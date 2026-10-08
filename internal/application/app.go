package application

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
)

type App struct {
	router *echo.Echo
	db     *pgx.Conn
	cfg    config
}

func NewApp(ctx context.Context, cfg config) *App {
	app := &App{
		cfg: cfg,
	}

	app.loadRoutes()
	app.connectToDatabase(ctx)

	return app
}

func (a *App) Start(ctx context.Context) {
	defer a.db.Close(ctx)

	if err := a.router.Start(":8080"); err != nil {
		a.router.Logger.Error("failed to start server", "error", err)
	}
}
