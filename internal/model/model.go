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

func (p URL) ResponseModel() URLResponse {
	return URLResponse{
		ID:          p.ID,
		URL:         p.URL,
		ShortCode:   p.ShortCode,
		CreatedAt:   p.CreatedAt.Format("January 2, 2006"),
		UpdatedAt:   p.UpdatedAt.Format("January 2, 2006"),
		AccessCount: p.AccessCount,
	}
}
