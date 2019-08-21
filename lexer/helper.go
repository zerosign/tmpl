package tests

import (
	"github.com/zerosign/tmpl/assert"
	"github.com/zerosign/tmpl/token"
	"testing"
)

// LexSpec: helper struct to define specs
//
type LexSpec struct {
	name, input string
	initial     flow.Flow
	expected    []token.Type
	fails       bool
}

// GroupedSpec: helper type that represents grouped LexSpec
//
type GroupedSpec map[string][]LexSpec

func RunGroupedSpec(testsets GroupedSpec, t *testing.T, name string) {
	if tests, ok := testsets[name]; ok {
		for _, test := range tests {
			if test.fails {
				// when test need to fails
				RunSpecFail(t, &test)
			} else {
				RunSpec(t, &test)
			}
		}
	} else {
		t.Errorf("no grouped test named %s", name)
	}
}

func RunSpecFail(t *testing.T, spec *LexSpec) {
	defer func() {
		if r := recover(); r != nil {
			t.Logf("spec fails is correct for %v", spec)
		}
	}()

	lexer := lexer.UnsafeNewLexer(spec.input, spec.initial)
	assert.AssertTokens(t, lexer, spec.expected)
}

func RunSpec(t *testing.T, spec *LexSpec) {
	lexer := lexer.UnsafeNewLexer(spec.input, spec.initial)
	assert.AssertTokens(t, lexer, spec.expected)
}
