package base

import (
	"fmt"
)

// Cursor: describe current position in text.
//
type Cursor struct {
	start, current int
}

// ZeroCursor: create cursor which points to start=0, current=0.
//
func ZeroCursor() Cursor {
	return Cursor{start: 0, current: 0}
}

// Start: Get the starting position of the cursor (caret)
//
//
func (c Cursor) Start() int {
	return c.start
}

// Current: Get current position of the cursor (caret).
//
//
func (c Cursor) Current() int {
	return c.current
}

// Next: increment cursor by 1.
//
//
func (c *Cursor) Next() {
	c.current += 1
}

// Incr: increment cursor by c.
//
func (c *Cursor) Incr(v int) {
	c.current += v
}

// Reset: reset the cursor to start=0, current=0.
//
func (c *Cursor) Reset() {
	c.start = 0
	c.current = 0
}

// Advance: advance the cursor start point to current point.
//
func (c *Cursor) Advance() {
	c.start = c.current
}

func (c Cursor) IsValid() bool {
	return c.current > c.start
}

// String: String representation for Cursor.
//
func (c Cursor) String() string {
	return fmt.Sprintf("<cursor{start=%d, current=%d}>", c.start, c.current)
}
