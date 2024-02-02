package hyprlang_parser

import (
	"errors"
)

type configFile struct {
	path    string   // Path of the file
	content []string // Content, line by line, of the file
	changed bool     // Store if config has changed
}

type Config struct {
	configFiles []configFile // List of config files (sources included)
	indentation int          // Indentation of the config file, based on the first section of the entry file
}

// LoadConfig loads config at {configPath} and returns a Config struct
// It also adds the sourced configuration files
func LoadConfig(configPath string) (config Config, err error) {
	err = config.addConfigFile(configPath)
	config.indentation = getIndentation(&config.configFiles[0].content)
	return
}

// WriteConfig writes/saves changed configs
func (c *Config) WriteConfig() error {
	for i, configFile := range c.configFiles {
		if configFile.changed {
			err := configFile.writeFile()
			if err != nil {
				return err
			}
			c.configFiles[i].changed = false
		}
	}
	return nil
}

// GetFirst returns the first value for {variable} at {section}
func (c *Config) GetFirst(section, variable string) string {
	var value string
	for _, configFile := range c.configFiles {
		values := getVariables(configFile.content, parseSectionString(section), variable, -1)
		if len(values) > 0 {
			return values[0]
		}
	}
	return value
}

// GetAll returns all the values for {variable} at {section}
func (c *Config) GetAll(section, variable string) []string {
	var values []string
	for _, configFile := range c.configFiles {
		values = append(values, getVariables(configFile.content, parseSectionString(section), variable, 0)...)
	}
	return values
}

// EditFirst changes the value of the first {variable} at {section} to {value}
// It returns an error if the variable was not found
func (c *Config) EditFirst(section, variable, value string) error {
	var isEdited bool
	for i, configFile := range c.configFiles {
		isEdited = editVariableN(&configFile.content, parseSectionString(section), variable, value, 0, c.indentation)
		if isEdited {
			c.configFiles[i].changed = true
			return nil
		}
	}

	return errors.New("Variable not found (Not edited).")
}

// EditN changes the value of the {n} {variable} at {section} to {value}
// It returns an error if the variable was not found
func (c *Config) EditN(section, variable, value string, n int) error {
	var isEdited bool
	for i, configFile := range c.configFiles {
		isEdited = editVariableN(&configFile.content, parseSectionString(section), variable, value, n, c.indentation)
		if isEdited {
			c.configFiles[i].changed = true
			return nil
		}
	}

	return errors.New("Variable not found (Not edited).")
}

// Add creates a {variable} at {section} with the {value} value
// It creates the sections if they don't exist
// If sections are not found, it will add the sections and the variable to the first config file
func (c *Config) Add(section, variable, value string) {
	var whereSectionExist *configFile
	var exist bool = false

	for _, configFile := range c.configFiles {
		exist = doesSectionExist(&configFile.content, parseSectionString(section))
		if exist {
			whereSectionExist = &configFile
			break
		}
	}

	if exist {
		addVariable(&whereSectionExist.content, parseSectionString(section), variable, value, c.indentation)
	} else {
		parsedSection := parseSectionString(section)
		for i := 0; i < len(parsedSection); i++ {
			if !doesSectionExist(&c.configFiles[0].content, parsedSection[:i+1]) {
				addSection(&c.configFiles[0].content, parsedSection[:i], parsedSection[i], c.indentation)
			}
		}
		addVariable(&c.configFiles[0].content, parsedSection, variable, value, c.indentation)
	}
}

// RemoveFirst removes the first {variable} at {section}
// It returns an error if the variable is not found
func (c *Config) RemoveFirst(section, variable string) error {
	var isRemoved bool
	for _, configFile := range c.configFiles {
		isRemoved = removeVariableN(&configFile.content, parseSectionString(section), variable, 0)
		if isRemoved {
			return nil
		}
	}
	return errors.New("Variable not found (Not removed).")
}

// RemoveN removes the {n} {variable} at {section}
// It returns an error if the variable is not found
func (c *Config) RemoveN(section, variable string, n int) error {
	var isRemoved bool
	for _, configFile := range c.configFiles {
		isRemoved = removeVariableN(&configFile.content, parseSectionString(section), variable, n)
		if isRemoved {
			return nil
		}
	}
	return errors.New("Variable not found (Not removed).")
}
