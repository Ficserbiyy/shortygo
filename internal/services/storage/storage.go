// Package storage defines the application storage.
package storage

import (
	"context"
	"fmt"

	"github.com/Ficserbiyy/shortygo/internal/model"
	"github.com/jackc/pgx/v5"
)

const (
	alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

func encodeBase62(num int) string {
	if num == 0 {
		return string(alphabet[0])
	}

	var arr []byte
	base := len(alphabet)

	for num > 0 {
		rem := num % base
		arr = append(arr, (alphabet[rem]))
		num /= base
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return string(arr)
}

func GetByShortcode(ctx context.Context, db *pgx.Conn, shortcode string) (model.URL, error) {
	var url model.URL

	err := db.QueryRow(ctx, getByShortcodeQuery, shortcode).Scan(
		&url.ID,
		&url.URL,
		&url.ShortCode,
		&url.CreatedAt,
		&url.UpdatedAt,
	)

	return url, err
}

func Create(ctx context.Context, db *pgx.Conn, data model.URL) (string, error) {
	var id int

	err := db.QueryRow(
		ctx,
		createQuery,
		data.URL,
	).Scan(&id)

	if err != nil {
		return "", fmt.Errorf("URL creation failed: %w", err)
	}

	shortcode := encodeBase62(id)
	_, err = db.Exec(ctx, updateShortcodeQuery, shortcode, id)

	return shortcode, err
}

func DeleteByShortcode(ctx context.Context, db *pgx.Conn, shortcode string) error {
	_, err := db.Exec(ctx, deleteByShortcodeQuery, shortcode)
	if err != nil {
		return fmt.Errorf("failed to remove URL %s: %w", shortcode, err)
	}

	return nil
}
