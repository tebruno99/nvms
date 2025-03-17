package scanner

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"time"
)

type MetaEvent struct {
	Path       string    `json:"path"`
	Size       int64     `json:"size"`
	ScannedAt  time.Time `json:"scannedAt"`
	ModifiedAt time.Time `json:"modifiedAt"`
	Error      error     `json:"error"`
}

type LocalScanner struct {
	opts *ScannerOption
}

// NewClient initializes Client by merging DefaultIdentityClientOptions with provided ClientOptions
func NewLocalScanner(opts ...Option[ScannerOption]) (*LocalScanner, error) {
	scr := &LocalScanner{opts: &ScannerOption{filter: NewExtensionFilter()}}

	for _, opt := range opts {
		opt(scr.opts)
	}

	return scr, nil
}

func (l LocalScanner) Scan(ctx context.Context) chan MetaEvent {
	resChan := make(chan MetaEvent)
	go func() {
		err := filepath.Walk(l.opts.path, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
				if !l.opts.filter.Filter(path) {
					me := MetaEvent{
						Path:       path,
						Size:       info.Size(),
						ScannedAt:  time.Now(),
						ModifiedAt: info.ModTime(),
						Error:      err,
					}
					resChan <- me
				}
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
		close(resChan)
	}()

	return resChan
}
