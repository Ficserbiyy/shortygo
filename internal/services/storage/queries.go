package storage

const (
	getByShortcodeQuery = `
		SELECT
			id,
			url,
			shortcode,
			created_at,
			updated_at
		FROM links
		WHERE id = $1;
	`
)
