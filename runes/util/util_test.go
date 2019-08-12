package util

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

func TestJoinRunes(t *testing.T) {
	texts := [][]rune{[]rune("hello"), []rune("world")}

	expected := "hello world"
	result := Join(texts, ' ')

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
