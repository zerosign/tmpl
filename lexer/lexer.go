package lexer

import (
	"fmt"
	"log"
	"unicode/utf8"
)

type Lexer struct {
	inner  []rune
	cursor Cursor
	state  StateFn
	tokens chan Token
}

// Taken from angstrom ocaml
// - skip_while
// - skip
// - take_while
// - advance
// - any_int

// NewLexer: create new lexer based on string.
//
// This method checks whether input string are valid utf8 string or not.
// If the string are invalid utf8 string, this method will returns error.
//
func NewLexer(input string) (*Lexer, error) {

	if !utf8.ValidString(input) {
		return nil, InvalidUtfInput()
	}

	lexer := Lexer{
		inner:  []rune(input),
		cursor: ZeroCursor(),
		state:  lexText,
		tokens: make(chan Token, 2),
	}

	return &lexer, nil
}

// Cursor: Get current cursor
//
func (l Lexer) Cursor() Cursor {
	return l.cursor
}

// CursorMut: Get mutable cursor referrence.
//
// Being used for `in-place` update.
//
func (l *Lexer) CursorMut() *Cursor {
	return &l.cursor
}

// HasNext: Check whether there is next state function or not
//
func (l Lexer) HasNext() bool {
	return l.state != nil
}

// Next: Get next token in the lexer queue.
//
func (l *Lexer) Next() (Token, error) {
	var err error
	var token Token

	if l.HasNext() {
		l.state, err = l.state(l)
		if err != nil {
			return Token{}, err
		}

		token = <-l.tokens
	}

	return token, nil
}

// Advance: Advance the cursor in lexer.
//
func (l *Lexer) Advance() {
	// beware cursor.Next ~ lexer.Advance
	l.cursor.Next()
}

// Available: Check availability of next character
//
func (l Lexer) Available() bool {
	return l.cursor.Current()+1 < len(l.inner)
}

func (l Lexer) Emit(t TokenType) {
	cursor := l.Cursor()

	// need l.current-1, since it's being prefixed with the next token
	value := l.inner[cursor.Start() : l.cursor.Current()-1]

	l.tokens <- Token{t, value}
	// assign l.start with l.current, means
	// the range (l.start..l.current) already being
	// consumed by the lexer
	l.start = l.current
}

func (l Lexer) RunesAhead() []rune {
	return l.inner[l.cursor.Current():]
}

func (l Lexer) CurrentRune() rune {
	return l.inner[l.cursor.Current()]
}

//
// Note: This advance current cursor
//
func (l Lexer) TakeWhile(conds ...RuneCond) {

	flag := true

	for {
		for _, cond := range conds {
			flag = flag && cond(l.CurrentRune())
		}

		if l.Available() && flag {
			l.Advance()
		} else {
			break
		}
	}
}

//
// Note: This advance current cursor
//
func (l Lexer) AcceptAll(conds ...RuneCond) bool {
	flag := true

	for _, cond := range conds {
		flag = flag && l.Accept(cond)
	}

	return flag
}

//
// Note: This advance current cursor
//
func (l Lexer) AcceptUntil(conds ...RuneCond) bool {
	flag := true

	for {
		flag = flag && l.AcceptAll(conds...)

		if flag && l.Available() {
			l.Advance()
		} else {
			break
		}
	}

	return flag
}

//
// Note: This advance current cursor
//
func (l Lexer) Accept(cond RuneCond) bool {
	if cond(l.CurrentRune()) {
		if l.Available() {
			l.Advance()
			return true
		}
	}

	// first rune is not an accepted character
	return false
}

func (l Lexer) Ignore(cond RuneCond) {
	for {
		ch := l.CurrentRune()
		if cond(ch) && l.Available() {
			l.Advance()
		} else {
			break
		}
	}
}

func (l Lexer) String() string {
	return fmt.Sprintf("<lexer{position=%s}>", l.Position())
}

func (l Lexer) Close() {
	log.Println("lexer is closed")
}
