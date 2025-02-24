package utils

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/must"
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

// GetTrimmedLines reads a file and returns its lines with leading and trailing whitespace removed.
func GetTrimmedLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, erero.Wro(err)
	}
	defer func() {
		must.Done(file.Close())
	}()

	var res []string
	var reader = bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n') // Read a line ending with '\n'
		if err != nil {
			if err == io.EOF {
				res = append(res, strings.TrimSpace(str)) // Append the last line. Never forget to do this.
				break
			}
			return nil, erero.Wro(err)
		}
		res = append(res, strings.TrimSpace(str))
	}
	return res, nil
}

// GetSubstringBetween extracts the substring between sSub and eSub. Exclude the sSub and eSub.
func GetSubstringBetween(s string, sSub, eSub string) string {
	if sIdx, eIdx := strings.Index(s, sSub), strings.LastIndex(s, eSub); sIdx >= 0 && eIdx >= 0 && eIdx >= sIdx+len(sSub) {
		return s[sIdx+len(sSub) : eIdx]
	}
	return "" // Return an empty string if no substring is found
}
