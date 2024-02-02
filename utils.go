package hyprlang_parser

import "strings"

func insert(slice *[]string, index int, value string) {
	if len(*slice) == index {
		*slice = append(*slice, value)
		return
	}
	*slice = append((*slice)[:index+1], (*slice)[index:]...)
	(*slice)[index] = value
}

func remove(slice *[]string, index int) {
	*slice = append((*slice)[:index], (*slice)[index+1:]...)
}

func parseSectionString(section string) []string {
	// / -> []
	// /general -> [general]
	// /general/ -> [general]
	// -> []
	// /general/input -> [general, input]
	section = strings.TrimPrefix(section, "/")
	section = strings.TrimSuffix(section, "/")
	if section == "" {
		return []string{}
	}
	return strings.Split(section, "/")
}
