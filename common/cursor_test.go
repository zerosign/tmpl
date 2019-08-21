package common

import (
	"testing"
)

func Assert(t *testing.T, message string, condition bool) {
	if !condition {
		t.Errorf(message)
	}
}

func TestZeroCursor(t *testing.T) {
	cursor := ZeroCursor()
	Assert(t, "cursor current should be 0", cursor.Current() == 0)
	Assert(t, "cursor start should be 0", cursor.Start() == 0)
}

func TestCursorIncr(t *testing.T) {
	cursor := ZeroCursor()
	cursor.Incr(1)
	Assert(t, "cursor current should be 1", cursor.Current() == 1)
	Assert(t, "cursor start should still be 0", cursor.Start() == 0)
}

func TestCursorAdvance(t *testing.T) {
	cursor := ZeroCursor()
	cursor.Advance()
	Assert(t, "cursor start should be the same as current", cursor.Start() == cursor.Current())
}

func TestCursorReset(t *testing.T) {
	cursor := ZeroCursor()
	cursor.Incr(1)
	cursor.Incr(2)

	cursor.Reset()

	Assert(t, "cursor should be equal to Cursor Zero", cursor == ZeroCursor())
}

func TestCursorIsValid(t *testing.T) {
	cursor := ZeroCursor()
	cursor.Incr(1)
	cursor.Incr(1)

	cursor.Advance()
	Assert(t, "cursor after advance should still be valid", cursor.IsValid())

	cursor.Reset()
	Assert(t, "cursor after reset should still be valid", cursor.IsValid())
}
