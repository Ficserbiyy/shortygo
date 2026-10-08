package application

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
)

// createTables initializes database tables.
func (a *App) createTables(ctx context.Context) error {
	// Create main table
	_, err := a.db.Exec(ctx, `
        CREATE TABLE IF NOT EXISTS links (
            id INTEGER GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
            url VARCHAR(100) NOT NULL,
			shortcode VARCHAR(30) NOT NULL,
            created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
			updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
        )
    `)

	if err != nil {
		return fmt.Errorf("unable to create links table: %w", err)
	}
	return nil
}

func (a *App) connectToDatabase(ctx context.Context) {
	conn, err := pgx.Connect(ctx, os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	// Ping the database
	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("unable to ping database: %v", err)
	}

	a.db = conn

	if err := a.createTables(ctx); err != nil {
		log.Fatal(err)
	}
}
