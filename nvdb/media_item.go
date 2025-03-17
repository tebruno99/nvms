package nvdb

import "time"

type MediaItem struct {
	ID             int64     `json:"id"`
	Path           string    `json:"path"`
	MimeType       string    `json:"mimeType"`
	FileModifiedAt time.Time `json:"modifiedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
