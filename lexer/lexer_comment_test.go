package tests

import (
	"github.com/zerosign/tmpl/assert"
	"github.com/zerosign/tmpl/lexer/tests"
	"github.com/zerosign/tmpl/token"
	"testing"
)

var (
	testsets tests.GroupedSpec
)

func createCommentSpec(name, input string) tests.LexSpec {
	return tests.LexSpec{
		name:     name,
		input:    input,
		initial:  flow.LexBlockCommentOpen,
		expected: []token.Type{token.TokenBlockCommentOpen, token.TokenBlockComment, token.TokenBlockCommentClose},
		fails:    false,
	}
}

func init() {
	testsets = map[string][]tests.LexSpec{
		"simple_comment": []tests.LexSpec{
			createCommentSpec("simple comment", "{# #}"),
			createCommentSpec("simple multiline comment", `{#
test hello world
#}`),
			createCommentSpec("block inside multiline block comment", "{# {{ test }} #}"),
			createCommentSpec("multiline block inside multiline block comment", `{# {{ test }} {# hehehe #} {{ for { key, value } in config("test") do }}
teste asahdioahsdoiah soidahsoid haoushdoiah oidha oishd
 daishd uoiahsoid haosih
{{ end }}
#}`),
		},
	}
}

func TestSimpleCommentSpec(t *testing.T)   { tests.RunGroupedSpec(testsets, t, "simple_comment") }
func TestMultipleCommentSpec(t *testing.T) { tests.RunGroupedSpec(testsets, t, "multiple_comment") }
