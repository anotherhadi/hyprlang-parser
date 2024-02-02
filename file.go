package hyprlang_parser

import (
	"bufio"
	"fmt"
	"os"
)

func readFile(filename string) (content []string, err error) {
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

func (cf *configFile) writeFile() error {
	file, err := os.OpenFile(cf.path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range cf.content {
		_, err := fmt.Fprintln(writer, line)
		if err != nil {
			return err
		}
	}

	return writer.Flush()
}

func (c *Config) addConfigFile(path string) error {
	var err error
	configFile := configFile{
		path: path,
	}
	configFile.content, err = readFile(path)
	if err != nil {
		return err
	}
	c.configFiles = append(c.configFiles, configFile)
	sources := c.GetAll("", "source")
	for _, source := range sources {
		alreadyAdded := false
		for _, conf := range c.configFiles {
			if source == conf.path {
				alreadyAdded = true
			}
		}
		if alreadyAdded {
			continue
		}
		err = c.addConfigFile(source)
		if err != nil {
			return err
		}
	}
	return nil
}
