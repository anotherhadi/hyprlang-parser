package utils

import (
	"strings"
)

// Remove white spaces & comments
func formatLine(line string) string {
	trimmedLine := strings.TrimSpace(line)
	if index := strings.Index(trimmedLine, "#"); index != -1 {
		trimmedLine = trimmedLine[:index]
	}

	return trimmedLine
}

func insert(slice []string, index int, value string) []string {
	if len(slice) == index {
		return append(slice, value)
	}
	slice = append(slice[:index+1], slice[index:]...)
	slice[index] = value
	return slice
}

func remove(slice []string, index int) []string {
	return append(slice[:index], slice[index+1:]...)
}

func sectionsEqual(section1, section2 []string) bool {
	if section1 == nil && section2 == nil || len(section1) == 0 && len(section2) == 0 {
		return true
	}
	if !(len(section1) == len(section2)) {
		return false
	}
	for i := 0; i < len(section1); i++ {
		if section1[i] != section2[i] {
			return false
		}
	}
	return true
}

// Is section2 inside of section1
func sectionInside(section1, section2 []string) bool {
	if len(section2) < len(section1) {
		return false
	}
	for i := 0; i < len(section1); i++ {
		if section1[i] != section2[i] {
			return false
		}
	}
	return true
}

func formatSection(section string) []string {
	if len(section) == 0 {
		return []string{}
	}
	section = strings.TrimSpace(section)
	section = strings.TrimPrefix(section, "/")
	section = strings.TrimSuffix(section, "/")
	if section == "" {
		return []string{}
	}
	return strings.Split(section, "/")
}

func formatValue(value string) string {
	value = strings.TrimSpace(value)
	return value
}

func doesSectionExist(content []string, section string) bool {
	var currentSection []string
	var searchSection []string = formatSection(section)
	for _, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
				if sectionsEqual(currentSection, searchSection) {
					return true
				}
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			}
		}
	}
	return false
}

// Get a section, return a list of result, with [0]:path/to/var [1]:value
func GetSection(content []string, section string) [][]string {
	var results [][]string
	var currentSection []string
	var searchSection []string = formatSection(section)
	for _, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			} else if sectionInside(searchSection, currentSection) {
				splitedTrimmedLine := strings.SplitN(trimmedLine, "=", 2)
				variable := strings.Join(currentSection, "/")
				if variable != "" {
					variable += "/"
				}
				variable += formatValue(splitedTrimmedLine[0])
				value := formatValue(splitedTrimmedLine[1])
				results = append(
					results,
					[]string{
						variable,
						value,
					},
				)
			}
		}
	}
	return results
}

// Return a list of values
func GetVariables(content []string, section, variable string) []string {
	var results []string
	var currentSection []string
	var searchSection []string = formatSection(section)
	for _, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			} else if sectionsEqual(currentSection, searchSection) {
				splitedTrimmedLine := strings.SplitN(trimmedLine, "=", 2)
				if formatValue(splitedTrimmedLine[0]) == variable {
					results = append(results, formatValue(splitedTrimmedLine[1]))
				}
			}
		}
	}
	return results
}

// Edit a variable, return the new content
func EditVariable(content []string, section, variable, value string, n int) []string {
	var i int
	var currentSection []string
	var searchSection []string = formatSection(section)
	for lineIndex, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			} else if sectionsEqual(currentSection, searchSection) {
				splitedTrimmedLine := strings.SplitN(trimmedLine, "=", 2)
				if formatValue(splitedTrimmedLine[0]) == variable {
					if i == n {
						content[lineIndex] = ""
						for j := 0; j < len(currentSection); j++ {
							content[lineIndex] += "  "
						}
						content[lineIndex] += variable + "=" + value
						return content
					}
					i++
				}
			}
		}
	}
	content = AddVariable(content, section, variable, value)
	return content
}

// Remove a variable, return the new content
func RemoveVariable(content []string, section, variable string, n int) []string {
	var i int
	var currentSection []string
	var searchSection []string = formatSection(section)
	for lineIndex, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			} else if sectionsEqual(currentSection, searchSection) {
				splitedTrimmedLine := strings.SplitN(trimmedLine, "=", 2)
				if formatValue(splitedTrimmedLine[0]) == variable {
					if i == n {
						content = remove(content, lineIndex)
						return content
					}
					i++
				}
			}
		}
	}
	return content
}

// Add a section, return the new content
func addSection(content []string, section, newSection string) []string {
	var currentSection []string
	var searchSection []string = formatSection(section)
	var formattedNewSection string = formatSection(newSection)[0]
	for lineNumber, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			} else if sectionsEqual(currentSection, searchSection) {
				space := ""
				for i := 0; i < len(currentSection); i++ {
					space += "  "
				}
				content = insert(content, lineNumber, space+formattedNewSection+"{")
				content = insert(content, lineNumber+1, space+"}")
				return content
			}
		}
	}
	return content
}

// Add a variable, return the new content
func AddVariable(content []string, section, variable, value string) []string {
	var currentSection []string
	var searchSection []string = formatSection(section)

	var currentSearchSection string
	for _, subSearchSection := range searchSection {
		sectionExist := doesSectionExist(content, currentSearchSection+subSearchSection)
		if !sectionExist {
			content = addSection(content, currentSearchSection, subSearchSection)
		}
		currentSearchSection += subSearchSection + "/"
	}

	for lineNumber, line := range content {
		trimmedLine := formatLine(line)
		if trimmedLine != "" {
			if len(searchSection) == 0 {
				content = insert(content, lineNumber, variable+"="+value)
				return content
			}
			if strings.HasSuffix(trimmedLine, "{") {
				currentSection = append(currentSection, strings.TrimSpace(strings.TrimSuffix(trimmedLine, "{")))
				if sectionsEqual(currentSection, searchSection) {
					space := ""
					for i := 0; i < len(currentSection); i++ {
						space += "  "
					}
					content = insert(content, lineNumber+1, space+variable+"="+value)
					return content
				}
			} else if strings.HasPrefix(trimmedLine, "}") {
				currentSection = currentSection[:len(currentSection)-1]
			}
		}
	}
	return content
}
