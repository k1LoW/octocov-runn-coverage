package runncoverage

import (
	"bufio"
	"io"
	"strings"

	"github.com/k1LoW/octocov/report"
)

// ParseMetadata parse stdin and return metadata
func ParseMetadata(in io.Reader) []*report.MetadataKV {
	mset := []*report.MetadataKV{}
	scanner := bufio.NewScanner(in)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || !strings.Contains(line, ":") {
			continue
		}
		splited := strings.SplitN(line, ":", 2)
		mset = append(mset, &report.MetadataKV{
			Key:   splited[0],
			Value: strings.TrimSpace(splited[1]),
		})
	}
	return mset
}
