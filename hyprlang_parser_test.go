package hyprlang_parser_test

import (
	"testing"

	"github.com/anotherhadi/hyprlang-parser"
)

const configPath = "test_config/example.conf"

func TestReadConfig(t *testing.T) {
	_, err := hyprlang_parser.ReadConfig(configPath)
	if err != nil {
		t.Error("Couldn't read the example.conf file")
	}
}

func TestGetFirst(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)
	var got string

	got = hyprlang_parser.GetFirst(content, "input/", "kb_layout")
	if got != "fr" {
		t.Error()
	}

	got = hyprlang_parser.GetFirst(content, "/input/", "kb_layout")
	if got != "fr" {
		t.Error()
	}

	got = hyprlang_parser.GetFirst(content, "/input", "kb_layout")
	if got != "fr" {
		t.Error()
	}

	got = hyprlang_parser.GetFirst(content, "/input", "kb_model")
	if got != "" {
		t.Error()
	}
}

func TestGetN(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)
	var got string

	got = hyprlang_parser.GetN(content, "input/", "kb_layout", 0)
	if got != "fr" {
		t.Error()
	}

	got = hyprlang_parser.GetN(content, "/input/", "kb_layout", 1)
	if got != "" {
		t.Error()
	}

	got = hyprlang_parser.GetN(content, "", "exec-once", 4)
	if got != "dunst" {
		t.Error()
	}

	got = hyprlang_parser.GetN(content, "/", "exec-once", 4)
	if got != "dunst" {
		t.Error()
	}
}

func TestGetAll(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)
	got := hyprlang_parser.GetAll(content, "", "monitor")
	if len(got) != 2 || got[0] != "eDP-1,2240x1400@60,0x0,1" || got[1] != "eDP-2,2240x1400@60,0x0,1" {
		t.Error()
	}

	got = hyprlang_parser.GetAll(content, "", "exec-once")
	if len(got) != 9 {
		t.Error()
	}
}

func TestGetSection(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)
	got := hyprlang_parser.GetSection(content, "misc")
	if len(got) != 1 {
		t.Error()
	}
	if got[0][0] != "misc/no_vfr" {
		t.Error()
	}
	if got[0][1] != "1" {
		t.Error()
	}
}

func TestAddVariable(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)

	newContent := hyprlang_parser.AddVariable(content, "input/", "keyboard", "is better than mouse")

	got := hyprlang_parser.GetFirst(newContent, "input", "keyboard")

	if got != "is better than mouse" {
		t.Error()
	}

	newContent = hyprlang_parser.AddVariable(content, "input/newSection", "keyboard", "is better than mouse")

	got = hyprlang_parser.GetFirst(newContent, "input/newSection", "keyboard")

	if got != "is better than mouse" {
		t.Error()
	}
}

func TestRemoveFirst(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)

	newContent := hyprlang_parser.RemoveFirst(content, "", "exec-once")

	got := hyprlang_parser.GetFirst(newContent, "", "exec-once")

	if got == "hyprlang_parser" {
		t.Error()
	}
}

func TestRemoveN(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)

	newContent := hyprlang_parser.RemoveN(content, "", "exec-once", 1)

	got := hyprlang_parser.GetN(newContent, "", "exec-once", 1)

	if got == "hyprlang_parser_2" {
		t.Error()
	}
}

func TestEditFirst(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)

	newContent := hyprlang_parser.EditFirst(content, "input", "kb_layout", "not-a-layout")

	got := hyprlang_parser.GetFirst(newContent, "input", "kb_layout")

	if got != "not-a-layout" {
		t.Error()
	}
}

func TestEditN(t *testing.T) {
	content, _ := hyprlang_parser.ReadConfig(configPath)

	newContent := hyprlang_parser.EditN(content, "/", "exec-once", "why don't u used yaml", 1)

	got := hyprlang_parser.GetN(newContent, "", "exec-once", 1)

	if got != "why don't u used yaml" {
		t.Error()
	}
}
