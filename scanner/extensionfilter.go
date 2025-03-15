package scanner

import "path/filepath"

type ExtensionFilter struct {
	filterList map[string]bool
}

func NewExtensionFilter(exts ...string) ExtensionFilter {
	filterList := make(map[string]bool, len(exts))
	for _, ext := range exts {
		filterList[ext] = true
	}

	return ExtensionFilter{filterList}
}

// Filter returns false if file should be processed. True if file shoudl be ignored
func (f ExtensionFilter) Filter(path string) bool {
	if path != "" {
		ext := filepath.Ext(path)
		if _, ok := f.filterList[ext]; ok {
			return false
		}
	}

	return true
}
