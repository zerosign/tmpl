package lexer

import (
	"fmt"
)

// InvalidUtfInput: return an error for invalid utf8 input
//
func InvalidUtfInput() error {
	return fmt.Errorf("invalid utf8 input")
}

func UnavailableFlow() error {
	return fmt.Errorf("flow is nil, hasNext returns false")
}

func InvalidCursor() error {
	return fmt.Errorf("no backward cursor available")
}

func LexerChannelClosed() error {
	return fmt.Errorf("channel is closed")
}
