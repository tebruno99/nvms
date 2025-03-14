package scanner

import (
	"fmt"
	"os"
	"path/filepath"
)

type LocalScanner struct {
	opts *ScannerOption
}

// NewClient initializes Client by merging DefaultIdentityClientOptions with provided ClientOptions
func NewLocalScanner(opts ...Option[ScannerOption]) (*LocalScanner, error) {
	scr := &LocalScanner{opts: &ScannerOption{}}

	for _, opt := range opts {
		opt(scr.opts)
	}

	return scr, nil
}

func (l LocalScanner) Scan() error {

	err := filepath.Walk(l.opts.path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			fmt.Printf("Found: %s\n", path)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
