package util

import (
	"fmt"
	"os"
	"strings"
)

type DayFile struct {
	content []byte
}

func LoadDayFile(day int) *DayFile {
	p := fmt.Sprintf("day%d.txt", day)
	content, err := os.ReadFile(p)
	if err != nil {
		panic(err)
	}

	return &DayFile{
		content: content,
	}
}

func (d *DayFile) Content() string {
	return string(d.content)
}

func (d *DayFile) split(line, c string) []string {
	parts := strings.Split(line, c)
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	cleanParts := make([]string, 0, len(parts))
	for _, parts := range parts {
		if parts != "" {
			cleanParts = append(cleanParts, parts)
		}
	}

	return cleanParts
}

func (d *DayFile) SplitMulti(a, b string) [][]string {
	lines := d.split(string(d.content), a)
	result := make([][]string, 0, len(lines))
	for _, line := range lines {
		parts := d.split(line, b)
		result = append(result, parts)
	}

	return result
}

func (d *DayFile) Split(c string) []string {
	return d.split(string(d.content), c)
}

func (d *DayFile) Lines() []string {
	return d.split(string(d.content), "\n")
}
