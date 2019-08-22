package runes

import (
	"testing"
)

func TestHasPrefixRunes(t *testing.T) {
	text := []rune("Hello world")
	prefix := []rune("Hello")

	if !HasPrefix(text, prefix) {
		t.Fatalf("'%s' should have prefix '%s'", string(text), string(prefix))
	}
}

func TestCompareRunes(t *testing.T) {
	var texts []rune = nil
	var expected []rune = nil

	if !Compare(texts, expected) {
		t.Fatalf("both text should be nil, texts: '%v', expected: '%v'", texts, expected)
	}

	texts = []rune("hello")
	expected = []rune("hello")

	if !Compare(texts, expected) {
		t.Fatalf("both text should be equal in elements, texts: '%v', expected: '%v'", texts, expected)
	}

	texts = []rune("hello ")
	expected = []rune("hello")

	if Compare(texts, expected) {
		t.Fatalf("both text shouldn't be equal in length, texts: '%v', expected: '%v'", texts, expected)
	}

	texts = []rune("hella")
	expected = []rune("hello")

	if Compare(texts, expected) {
		t.Fatalf("both text shouldn't be equal in elements, texts: '%v', expected: '%v'", texts, expected)
	}
}

func TestJoinRunes(t *testing.T) {
	texts := [][]rune{[]rune("hello"), []rune("world")}

	expected := []rune("hello world")

	result := JoinRunes(texts, ' ')

	if !Compare(result, expected) {
		t.Fatalf("'%v' when joined should be '%v' but got '%v'", texts, expected, result)
	}

}

func TestJoinString(t *testing.T) {
	texts := [][]rune{[]rune("hello"), []rune("world")}

	expected := "hello world"
	result := JoinString(texts, ' ')

	if result != expected {
		t.Fatalf("'%v' when joined should be '%v' but got '%v'", texts, expected, result)
	}

}

func TestAnyPrefixesRunes(t *testing.T) {
	text := []rune("Hello world")
	prefixes := [][]rune{
		[]rune("H"),
		[]rune("He"),
		[]rune("Hello"),
	}

	if !AnyPrefixes(text, prefixes) {
		t.Fatalf("'%s' should have at least 1 prefix '%v'", string(text), prefixes)
	}
}
