package utils

import (
	"runtime"
	"strings"
)

type StringBuilder struct {
	lines []string
}

func NewStringBuilder() *StringBuilder {
	return &StringBuilder{
		lines: []string{},
	}
}

func (builder *StringBuilder) AppendLine(line string) {
	builder.lines = append(builder.lines, line)
}

func (builder *StringBuilder) String() string {
	newline := "\n"
	if runtime.GOOS == "windows" {
		newline = "\r\n"
	}

	return strings.Join(builder.lines, newline)
}
