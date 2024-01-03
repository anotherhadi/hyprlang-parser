package hyprlang_parser

import (
	"bufio"
	"fmt"
	"os"

	"github.com/anotherhadi/hyprlang_parser/utils"
)

func ReadConfig(filename string) (content []string, err error) {
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		content = append(content, line)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return content, nil
}

func WriteConfig(content []string, filename string) (err error) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range content {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func GetFirst(content []string, section, variable string) string {
	variables := utils.GetVariables(content, section, variable)
	if len(variables) == 0 {
		return ""
	}
	return variables[0]
}

func GetN(content []string, section, variable string, n int) string {
	variables := utils.GetVariables(content, section, variable)
	if n > len(variables)-1 {
		return ""
	}
	return variables[n]
}

func GetAll(content []string, section, variable string) []string {
	return utils.GetVariables(content, section, variable)
}

func GetSection(content []string, section string) [][]string {
	return utils.GetSection(content, section)
}

func AddVariable(content []string, section, variable, value string) []string {
	return utils.AddVariable(content, section, variable, value)
}

func RemoveFirst(content []string, section, variable string) []string {
	return utils.RemoveVariable(content, section, variable, 0)
}

func RemoveN(content []string, section, variable string, n int) []string {
	return utils.RemoveVariable(content, section, variable, n)
}

func EditFirst(content []string, section, variable, value string) []string {
	return utils.EditVariable(content, section, variable, value, 0)
}

func EditN(content []string, section, variable, value string, n int) []string {
	return utils.EditVariable(content, section, variable, value, n)
}
