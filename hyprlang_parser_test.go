package hyprlang_parser_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/anotherhadi/hyprlang-parser"
)

const configPath = "test_config/example.conf"

func strShouldEqual(t *testing.T, str1 string, str2 string) {
	if str1 != str2 {
		t.Error("Strings '" + str1 + "' and '" + str2 + "' not equal")
	}
}

func strsShouldEqual(t *testing.T, strs1 []string, strs2 []string) {
	if !reflect.DeepEqual(strs1, strs2) {
		t.Error("[]Strings '" + fmt.Sprint(strs1) + "' and '" + fmt.Sprint(strs2) + "' not equal")
	}
}

func failOnError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func TestLoadConfig(t *testing.T) {
	_, err := hyprlang_parser.LoadConfig(configPath)
	if err != nil {
		t.Error(err)
	}
}

func TestWriteConfig(t *testing.T) {
	config, err := hyprlang_parser.LoadConfig(configPath)
	if err != nil {
		t.Error(err)
	}
	err = config.WriteConfig()
	if err != nil {
		t.Error(err)
	}
}

func TestGetFirst(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetFirst("", "exec-once"), "hyprlang_parser")
	strShouldEqual(t, config.GetFirst("", "exec_once"), "")
	strShouldEqual(t, config.GetFirst("input", "kb_layout"), "fr")
	strShouldEqual(t, config.GetFirst("input/", "kb_layout"), "fr")
	strShouldEqual(t, config.GetFirst("/input/", "kb_layout"), "fr")
	strShouldEqual(t, config.GetFirst("", "kb_layout"), "")
	strShouldEqual(t, config.GetFirst("input/touchpad/", "natural_scroll"), "yes")
	strShouldEqual(t, config.GetFirst("/input/touchpad", "natural_scroll"), "yes")

	// Sourced files
	strShouldEqual(t, config.GetFirst("", "only-sourced-var"), "ok")
	strShouldEqual(t, config.GetFirst("section/subsection", "does-it-work"), "true")
}

func TestGetAll(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strsShouldEqual(t, config.GetAll("", "monitor"), []string{"eDP-1,2240x1400@60,0x0,1", "eDP-2,2240x1400@60,0x0,1"})
}

func TestEditFirst(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetFirst("input", "kb_layout"), "fr")
	failOnError(t, config.EditFirst("input/", "kb_layout", "en"))
	strShouldEqual(t, config.GetFirst("input", "kb_layout"), "en")

	strShouldEqual(t, config.GetFirst("input/touchpad", "natural_scroll"), "yes")
	failOnError(t, config.EditFirst("input/touchpad", "natural_scroll", "no"))
	strShouldEqual(t, config.GetFirst("input/touchpad", "natural_scroll"), "no")

	strShouldEqual(t, config.GetFirst("", "monitor"), "eDP-1,2240x1400@60,0x0,1")
	failOnError(t, config.EditFirst("", "monitor", "no monitor here"))
	strShouldEqual(t, config.GetFirst("", "monitor"), "no monitor here")

	strShouldEqual(t, config.GetFirst("", "only-sourced-var"), "ok")
	failOnError(t, config.EditFirst("", "only-sourced-var", "youpi"))
	strShouldEqual(t, config.GetFirst("", "only-sourced-var"), "youpi")
}

func TestEditN(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetAll("", "exec-once")[1], "hyprlang_parser_2")
	failOnError(t, config.EditN("/", "exec-once", "hyprsettings", 1))
	strShouldEqual(t, config.GetAll("", "exec-once")[1], "hyprsettings")
}

func TestAdd(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetFirst("", "a-new-var"), "")
	config.Add("", "a-new-var", "newvalue")
	strShouldEqual(t, config.GetFirst("", "a-new-var"), "newvalue")

	strShouldEqual(t, config.GetFirst("test", "a-new-var"), "")
	config.Add("test", "a-new-var", "newvalue")
	strShouldEqual(t, config.GetFirst("test", "a-new-var"), "newvalue")

	strShouldEqual(t, config.GetFirst("test/subsections", "a-new-var"), "")
	config.Add("test/subsections", "a-new-var", "newvalue")
	strShouldEqual(t, config.GetFirst("test/subsections", "a-new-var"), "newvalue")
}

func TestRemoveFirst(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetFirst("", "monitor"), "eDP-1,2240x1400@60,0x0,1")
	failOnError(t, config.RemoveFirst("", "monitor"))
	strShouldEqual(t, config.GetFirst("", "monitor"), "eDP-2,2240x1400@60,0x0,1")

	failOnError(t, config.RemoveFirst("animations", "enabled"))
	strShouldEqual(t, config.GetFirst("animations", "enabled"), "")
}

func TestRemoveN(t *testing.T) {
	config, _ := hyprlang_parser.LoadConfig(configPath)

	strShouldEqual(t, config.GetAll("animations", "animation")[1], "border,1,10,default")
	failOnError(t, config.RemoveN("animations", "animation", 1))
	strShouldEqual(t, config.GetAll("animations", "animation")[1], "fade,0,5,default")

}
