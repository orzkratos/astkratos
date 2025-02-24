package astkratos

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
)

// SuffixMatcher checks if a string ends with any of the specified suffixes.
type SuffixMatcher struct {
	suffixes []string
}

// NewSuffixMatcher creates a new SuffixMatcher.
func NewSuffixMatcher(suffixes []string) *SuffixMatcher {
	return &SuffixMatcher{
		suffixes: suffixes,
	}
}

// Match checks if the string ends with any of the suffixes.
func (sm *SuffixMatcher) Match(s string) bool {
	for _, suffix := range sm.suffixes {
		if strings.HasSuffix(s, suffix) {
			return true
		}
	}
	return false
}

// WalkFiles traverses the file tree rooted at root, applying the run function to each file that matches the suffixes.
func WalkFiles(root string, suffixMatcher *SuffixMatcher, run func(path string, info os.FileInfo) error) error {
	if err := filepath.Walk(root,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return erero.Wro(err)
			}
			if info == nil || info.IsDir() {
				return nil
			}
			if suffixMatcher.Match(path) {
				return run(path, info)
			}
			return nil
		},
	); err != nil {
		return erero.Wro(err)
	}
	return nil
}
