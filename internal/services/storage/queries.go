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
		WHERE shortcode = $1;
	`

	createQuery = `
		INSERT INTO expenses (url, shortcode)
        VALUES ($1, $2)
        RETURNING id
	`

	updateShortcodeQuery = `
		UPDATE links
		SET shortcode = $1
		WHERE id = $2
	`

	deleteByShortcodeQuery = `
		DELETE FROM links
		WHERE shortcode = $1
	`
)
