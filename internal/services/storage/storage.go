package storage

import (
	"context"
	"fmt"

	"github.com/Ficserbiyy/shortygo/internal/model"
	"github.com/jackc/pgx/v5"
)

func GetURLByShortcode(ctx context.Context, db *pgx.Conn, shortcode string) (model.URL, error) {
	var url model.URL

	err := db.QueryRow(ctx, getByShortcodeQuery, shortcode).Scan(
		&url.ID,
		&url.URL,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)

	if err != nil {
		return model.URL{}, fmt.Errorf("failed to query url: %w", err)
	}
	return url, nil
}
