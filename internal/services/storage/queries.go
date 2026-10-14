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
		INSERT INTO links (url)
        VALUES ($1)
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
