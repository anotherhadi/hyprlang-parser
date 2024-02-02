package hyprlang_parser

import (
	"reflect"
	"strings"
)

// Return only one value if n == -1
func getVariables(content, section []string, variable string, n int) []string {
	var values []string
	var currentSection []string = []string{}
	for _, line := range content {
		if skipLine(line) {
			continue
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		} else if currentLineType == lineVariable {
			if !reflect.DeepEqual(currentSection, section) {
				continue
			}
			name, value := parseLineVariable(line)
			if name == variable {
				values = append(values, value)
				if n == -1 {
					return values
				}
			}
		}
	}
	return values
}

func editVariableN(content *[]string, section []string, variable string, newValue string, n int, indentation int) bool {
	var currentSection []string = []string{}
	var i int
	for lineNumber, line := range *content {
		if skipLine(line) {
			continue
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		} else if currentLineType == lineVariable {
			if !reflect.DeepEqual(currentSection, section) {
				continue
			}
			name, _ := parseLineVariable(line)
			if name == variable {
				if i == n {
					(*content)[lineNumber] = strings.Repeat(" ", indentation*len(currentSection)) + name + "=" + newValue
					return true
				}
				i++
			}
		}
	}
	return false
}

func removeVariableN(content *[]string, section []string, variable string, n int) bool {
	var currentSection []string = []string{}
	var i int
	for lineNumber, line := range *content {
		if skipLine(line) {
			continue
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		} else if currentLineType == lineVariable {
			if !reflect.DeepEqual(currentSection, section) {
				continue
			}
			name, _ := parseLineVariable(line)
			if name == variable {
				if i == n {
					remove(content, lineNumber)
					return true
				}
				i++
			}
		}
	}
	return false
}

func doesSectionExist(content *[]string, section []string) bool {
	if len(section) == 0 {
		return true
	}
	var currentSection []string = []string{}
	for _, line := range *content {
		if skipLine(line) {
			continue
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
			if reflect.DeepEqual(currentSection, section) {
				return true
			}
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		}
	}
	return false
}

func getIndentation(content *[]string) int {
	for lineNumber, line := range *content {
		if skipLine(line) {
			continue
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			if lineNumber < len(*content) {
				count := 0
				for _, char := range (*content)[lineNumber+1] {
					if char == ' ' {
						count++
					} else {
						break
					}
				}
				return count
			}
		}
	}
	return 2
}

func addSection(content *[]string, section []string, newSection string, indentation int) {
	var currentSection []string = []string{}
	for lineNumber, line := range *content {
		if skipLine(line) {
			continue
		}
		if reflect.DeepEqual(currentSection, section) {
			insert(content, lineNumber, strings.Repeat(" ", indentation*len(currentSection))+newSection+" {")
			insert(content, lineNumber+1, strings.Repeat(" ", indentation*len(currentSection))+"}")
			return
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		}
	}
}

func addVariable(content *[]string, section []string, variable string, value string, indentation int) {
	var currentSection []string = []string{}
	for lineNumber, line := range *content {
		if skipLine(line) {
			continue
		}
		if reflect.DeepEqual(currentSection, section) {
			insert(content, lineNumber, strings.Repeat(" ", (len(currentSection)+1)*indentation)+variable+"="+value)
			return
		}
		currentLineType := getLineType(line)
		if currentLineType == lineSectionStart {
			currentSection = append(currentSection, parseLineSectionStart(line))
		} else if currentLineType == lineSectionEnd {
			if len(currentSection) > 0 {
				currentSection = currentSection[:len(currentSection)-1]
			}
		}
	}
}
