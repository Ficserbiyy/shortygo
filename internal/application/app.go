package application

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v5"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router *echo.Echo
	db     *pgx.Conn
	redis  *redis.Client
	cfg    config
}

func NewApp(ctx context.Context, cfg config) *App {
	app := &App{
		redis: redis.NewClient(&redis.Options{
			Addr: cfg.redisAddress,
		}),
		cfg: cfg,
	}

	app.connectToDatabase(ctx)
	app.loadRoutes()

	return app
}

func (a *App) Start(ctx context.Context) {
	defer a.db.Close(ctx)

	srvPort := fmt.Sprintf(":%d", a.cfg.serverPort)

	if err := a.router.Start(srvPort); err != nil {
		a.router.Logger.Error("failed to start server", "error", err)
	}
}
