package lexer

import (
	"github.com/zerosign/tmpl/token"
	"testing"
)

// UnsafeDefaultLexer: Utility function for creating new lexer with default flow.
//
// will raise fatal error if fails
//
func UnsafeDefaultLexer(input string) Lexer {
	nlexer, err := DefaultLexer(input)

	if err != nil {
		panic(err)
	}

	return nlexer
}

// UnsafeNewLexer: Utility function for creating new lexer with given flow.
//
func UnsafeNewLexer(input string, f Flow) Lexer {
	nlexer, err := NewLexer(input, f)

	if err != nil {
		panic(err)
	}

	return nlexer
}

// AssertNextToken: Utility function for fetching next token
//
// will raise fatal error if fails
//
func AssertNextToken(t *testing.T, lexer Lexer) *token.Token {
	token, err := lexer.Next()
	if err != nil {
		t.Fatal(err)
	}

	return &token
}

// AssertToken: Utility function for checking for expected token type.
//
// will raise fatal error if fails
//
func AssertToken(t *testing.T, result *token.Token, expected token.Type) {
	if result.Type != expected {
		t.Fatalf("token should be %s but got %s", expected.String(), result)
	}
}

// AssertTokens: Utility function for checking for list of expected token type in order.
//
// will raise fatal error if one of expected token aren't equal (in order).
//
func AssertTokens(t *testing.T, lexer Lexer, expectedTypes []token.Type) {
	var token *token.Token
	for ii := 0; ii < len(expectedTypes); ii += 1 {
		if lexer.HasNext() {
			token = AssertNextToken(t, lexer)
			t.Logf("asserted_token: %v", token)
			AssertToken(t, token, expectedTypes[ii])
		} else {
			t.Fatalf("lexer should have at least %d tokens left but got 0", len(expectedTypes)-ii)
		}
	}
}

// LexSpec: helper struct to define specs
//
type LexSpec struct {
	name, input string
	initial     Flow
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

	lexer := UnsafeNewLexer(spec.input, spec.initial)
	AssertTokens(t, lexer, spec.expected)
}

func RunSpec(t *testing.T, spec *LexSpec) {
	lexer := UnsafeNewLexer(spec.input, spec.initial)
	AssertTokens(t, lexer, spec.expected)
}
