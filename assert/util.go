package assert

import (
	"github.com/zerosign/tmpl/token"
	"testing"
)

// AssertNextToken: Utility function for fetching next token
//
// will raise fatal error if fails
//
func AssertNextToken(t *testing.T, lexer base.Lexer) *token.Token {
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
func AssertTokens(t *testing.T, lexer base.Lexer, expectedTypes []token.Type) {
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
