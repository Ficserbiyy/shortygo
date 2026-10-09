package model

import "time"

type URL struct {
	ID          uint64    `json:"id,omitempty"`
	URL         string    `json:"url,omitempty"`
	ShortCode   string    `json:"short_code,omitempty"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
	AccessCount uint      `json:"access_count,omitempty"`
}

type URLResponse struct {
	ID          uint64 `json:"id,omitempty"`
	URL         string `json:"url,omitempty"`
	ShortCode   string `json:"short_code,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	UpdatedAt   string `json:"updated_at,omitempty"`
	AccessCount uint   `json:"access_count,omitempty"`
}
