# Hyprlang Parser

<p>
    <a href="https://github.com/anotherhadi/hyprlang-parser/releases"><img src="https://img.shields.io/github/release/anotherhadi/md-table-of-contents.svg" alt="Latest Release"></a>
    <a href="https://pkg.go.dev/github.com/anotherhadi/hyprlang-parser?tab=doc"><img src="https://godoc.org/github.com/golang/gddo?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/anotherhadi/hyprlang-parser"><img src="https://goreportcard.com/badge/github.com/anotherhadi/md-table-of-contents" alt="GoReportCard"></a>
</p>

A Golang implementation library for the hypr config language.

## Functions

**Understanding Section and Variable Parameters:**

For most functions, you'll need to provide a "section" and a "variable" parameter.

- The "section" represents the name of the section, which is separated by forward slashes (/).
- The main section can be represented by either an empty string ("") or a single forward slash ("/").
- A decoration section could be "decoration" (or "/decoration/", as per your preference).

So, for example, to retrieve the enabled variable in the decoration section for blur:

- Set `section="decoration/blur"`
- Set `variable="enabled"`

---

| Function         | Explanation                                    | Example                                         |
|------------------|------------------------------------------------|-------------------------------------------------|
| LoadConfig       | Loads a configuration file at the specified path and returns a Config struct. It also adds the sourced configuration files. | ```config, err := LoadConfig("~/.config/hypr/hyprland.conf")``` |
| WriteConfig      | Writes/saves changed configurations.          | ```err := config.WriteConfig()```            |
| GetFirst         | Returns the first value for the specified variable in the given section. | ```value := config.GetFirst("input", "kb_layout")``` |
| GetAll           | Returns all the values for the specified variable in the given section. | ```values := config.GetAll("", "exec-once")``` |
| EditFirst        | Changes the value of the first occurrence of the specified variable in the given section to the provided value. Returns an error if the variable was not found. | ```err := config.EditFirst("input/touchpad", "natural_scroll", "true")``` |
| EditN            | Changes the value of the nth occurrence of the specified variable in the given section to the provided value. Returns an error if the variable was not found. | ```err := config.EditN("animations", "bezier", "winIn, 0.05, 0.9, 0.1, 1.1", 2)``` |
| Add              | Creates a new variable with the provided value in the specified section. It creates the sections if they don't exist. If sections are not found, it will add the sections and the variable to the first config file. | ```config.Add("decoration/blur", "size", "3")``` |
| RemoveFirst      | Removes the first occurrence of the specified variable in the given section. Returns an error if the variable is not found. | ```err := config.RemoveFirst("", "exec-once")``` |
| RemoveN          | Removes the nth occurrence of the specified variable in the given section. Returns an error if the variable is not found. | ```err := config.RemoveN("", "exec-once", 2)``` |

## Example:

You can find more examples and usage in [`hyprlang_parser_test.go`](hyprland_parser_test.go).
