package tests

import (
	// "github.com/rs/zerolog/log"
	"github.com/zerosign/tmpl/lexer/flow"
	"github.com/zerosign/tmpl/token"
	"testing"
)

var (
	testsets GroupedSpec
)

func init() {
	testsets = map[string][]LexSpec{
		"number": []LexSpec{
			LexSpec{
				name:     "simple number",
				input:    "1",
				initial:  flow.LexNumber,
				expected: []token.Type{token.Integer},
				fails:    false,
			},
			LexSpec{
				name:     "zero padded number",
				input:    "001",
				initial:  flow.LexNumber,
				expected: []token.Type{},
				fails:    true,
			},
			LexSpec{
				name:     "simple recurrent number",
				input:    "12121312",
				initial:  flow.LexNumber,
				expected: []token.Type{token.Integer},
				fails:    false,
			},
			LexSpec{
				name:     "simple double",
				input:    "1.2",
				initial:  flow.LexNumber,
				expected: []token.Type{token.Double},
				fails:    false,
			},
			LexSpec{
				name:     "long double",
				input:    "0.00000000000000001",
				initial:  flow.LexNumber,
				expected: []token.Type{token.Double},
				failse:   false,
			},
			LexSpec{
				name:     "correct double with 0",
				input:    "0.1001",
				initial:  flow.LexNumber,
				expected: []token.Type{},
				fails:    false,
			},
		},
		"string": []LexSpec{
			LexSpec{
				name:     "empty string",
				input:    "\"\"",
				initial:  flow.LexString,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "whitespace string",
				input:    "\n\r\t    \n   ",
				initial:  flow.LexString,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				input:    "Hello world!",
				initial:  flow.LexString,
				expected: []token.Type{},
				fails:    true,
			},
			LexSpec{
				input:    "\"Hello world!\"",
				initial:  flow.LexString,
				expected: []token.Type{token.String},
				fails:    false,
			},
			LexSpec{
				input:    "\"1212 w1wdqwqwqweqw\"",
				initial:  flow.LexString,
				expected: []token.Type{token.String},
				fails:    false,
			},
		},
		"map": []LexSpec{
			LexSpec{
				name:     "empty map",
				input:    "{}",
				initial:  flow.LexMap,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "simple map object",
				input:    "{\"test\": 1}",
				initial:  flow.LexMap,
				expected: []token.Type{},
				fails:    false,
			},
		},
		"array": []LexSpec{
			LexSpec{
				name:     "empty array",
				input:    "[]",
				initial:  flow.LexArray,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "string array",
				input:    "[\"test\", \"hello world\"]",
				initial:  flow.LexArray,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "integer array",
				input:    "[1, 2, 3, 4, 5, 6, 10]",
				initial:  flow.LexArray,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "double array",
				input:    "[2.0, 0.0, 0.1]",
				initial:  flow.LexArray,
				expected: []token.Type{},
				fails:    false,
			},
		},
		"value": []LexSpec{
			LexSpec{
				name:     "integer value",
				input:    "1",
				initial:  flow.LexValue,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "double value",
				input:    "1.1",
				initial:  flow.LexValue,
				expected: []token.Type{},
				fails:    false,
			},
			LexSpec{
				name:     "string value",
				input:    "\"hello world\"",
				initial:  flow.LexValue,
				expected: []token.Type{},
				fails:    false,
			},
		},
	}
}

func TestNumberLexer(t *testing.T) { RunGroupedSpec(testsets, t, "number") }
func TestStringLexer(t *testing.T) { RunGroupedSpec(testsets, t, "string") }
func TestArrayLexer(t *testing.T)  { RunGroupedSpec(testsets, t, "array") }
func TestMapLexer(t *testing.T)    { RunGroupedSpec(testsets, t, "map") }
func TestValueLexer(t *testing.T)  { RunGroupedSpec(testsets, t, "value") }
