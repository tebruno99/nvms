package scanner

import "time"

type MediaScanner interface {
	Scan() error
}

type Filter interface {
	Filter(path string) bool
}

type MetaEvent struct {
	Path       string    `json:"path"`
	Size       int64     `json:"size"`
	ScannedAt  time.Time `json:"scannedAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Error      error     `json:"error"`
}
