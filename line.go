package hyprlang_parser

import (
	"strings"
)

func formatLine(line string) string {
	line = strings.TrimSpace(line)
	if !strings.Contains(line, "#") {
		return line
	}
	for i, letter := range line {
		if letter == '#' {
			if i-1 >= 0 && line[i-1] == '#' {
				continue
			}
			if i+1 < len(line)-1 && line[i+1] == '#' {
				continue
			}
			return line[:i]
		}
	}
	return line
}

func skipLine(line string) bool {
	line = formatLine(line)
	if line == "" {
		return true
	}
	if strings.HasPrefix(line, "#") {
		return true
	}

	return false
}

type lineType int

const (
	lineSectionStart lineType = 0
	lineSectionEnd   lineType = 1
	lineVariable     lineType = 2
)

func getLineType(line string) lineType {
	line = formatLine(line)
	if strings.HasSuffix(line, "{") {
		return lineSectionStart
	}
	if strings.HasPrefix(line, "}") {
		return lineSectionEnd
	}
	return lineVariable
}

func parseLineVariable(line string) (name, value string) {
	line = formatLine(line)
	lineSplitted := strings.SplitN(line, "=", 2)
	name = strings.TrimSpace(lineSplitted[0])
	value = strings.TrimSpace(lineSplitted[1])
	return
}

func parseLineSectionStart(line string) (sectionName string) {
	line = formatLine(line)
	sectionName = strings.TrimSuffix(line, "{")
	sectionName = strings.TrimSpace(sectionName)
	return
}
