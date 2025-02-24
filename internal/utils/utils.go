package utils

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/yyle88/erero"
	"github.com/yyle88/must"
)

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
