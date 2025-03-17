package nvdb

import (
	"time"
)

type File struct {
	ID             int64     `json:"id`
	Name           string    `json:"name"`
	Extension      string    `json:"extension"`
	BasePath       string    `json:"basePath"`
	Path           string    `json:"path"`
	Size           int64     `json:"size"`
	LibraryPathID  int64     `json:"libraryPathID"`
	FileModifiedAt time.Time `json:"fileModifiedAt"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
