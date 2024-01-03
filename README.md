# Hyprlang Parser

A Golang implementation library for the hypr config language.

## Example:

```go
content, err := hyprlang_parser.ReadConfig(configPath)
if err != nil {
  panic(err)
}

var result string

result = hyprlang_parser.GetFirst(content, "input/", "kb_layout")
fmt.Println(result) // fr

content = hyprlang_parser.EditFirst(content, "/input", "kb_layout", "it")

result = hyprlang_parser.GetFirst(content, "input/", "kb_layout")
fmt.Println(result) // it
```

You can found more example/usage on the `hyprlang_parser_test.go` file
