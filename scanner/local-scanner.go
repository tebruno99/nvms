package scanner

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

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

func (l LocalScanner) Scan() error {
	err := filepath.Walk(l.opts.path, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			if !l.opts.filter.Filter(path) {
				f := MetaEvent{
					Path:       path,
					Size:       info.Size(),
					ModifiedAt: info.ModTime(),
				}
				bt, _ := json.Marshal(f)
				fmt.Printf("Found: %s\n", bt)
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
