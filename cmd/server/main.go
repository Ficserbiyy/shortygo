package main

import (
	"context"

	"github.com/Ficserbiyy/shortygo/internal/application"
)

func main() {
	ctx := context.Background()

	app := application.NewApp(ctx, application.LoadConfig())

	app.Start(ctx)
}
