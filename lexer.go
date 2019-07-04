package tmpl

import (
	"fmt"
	"log"
	"unicode/utf8"
)

type Position struct {
	start, current int
}

func (p Position) String() string {
	return fmt.Sprintf("<position{start=%d, current=%d}>", p.start, p.current)
}

type Lexer struct {
	inner          []rune
	start, current int
	tokens         chan Token
}

// Taken from angstrom ocaml
// - skip_while
// - skip
// - take_while
// - advance
// - any_int

func NewLexer(input string) (*Lexer, chan Token, error) {
	if !utf8.ValidString(input) {
		return nil, nil, InvalidUtfInput()
	}

	lexer := Lexer{
		inner: []rune(input),
		start: 0, current: 0,
		tokens: make(chan Token, 2),
	}

	go lexer.Run()

	return &lexer, lexer.tokens, nil
}

func (l Lexer) Position() Position {
	return Position{l.start, l.current}
}

func (l Lexer) Run() {
	var nextStateFn StateFn
	var err error

	// close tokens channel
	defer close(l.tokens)
	nextStateFn, err = lexText(&l)

	for {
		if err != nil {
			log.Print(err)
			return
		}

		if nextStateFn != nil {
			nextStateFn, err = nextStateFn(&l)
		}
	}
}

//
// Advance
//
func (l *Lexer) Advance() {
	l.current = l.current + 1
}

func (l Lexer) HasNext() bool {
	return l.current+1 < len(l.inner)
}

func (l Lexer) Emit(t TokenType) {
	// need l.current-1, since it's being prefixed with the next token
	log.Printf("start: %d, current: %d, length: %d", l.start, l.current-1, len(l.inner))
	value := l.inner[l.start : l.current-1]

	l.tokens <- Token{t, value}
	// assign l.start with l.current, means
	// the range (l.start..l.current) already being
	// consumed by the lexer
	l.start = l.current
}

func (l Lexer) RunesAhead() []rune {
	return l.inner[l.current:]
}

func (l Lexer) CurrentRune() rune {
	return l.inner[l.current]
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

		if l.HasNext() && flag {
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

		if flag && l.HasNext() {
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
		if l.HasNext() {
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
		if cond(ch) && l.HasNext() {
			l.Advance()
		} else {
			break
		}
	}
}

func (l Lexer) String() string {
	return fmt.Sprintf("<lexer{position=%s}>", l.Position())
}
