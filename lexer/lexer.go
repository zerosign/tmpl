package lexer

import (
	"fmt"
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/lexer/flow"
	"github.com/zerosign/tmpl/token"
	"log"
	"unicode/utf8"
)

// GenLexer : types that implement lexer.Lexer for this language
//
//
type GenLexer struct {
	inner  []rune
	cursor base.Cursor
	flow   flow.Flow
	tokens chan token.Token
}

// NewLexer: create new lexer based on string by specifying initial state.
//
// This method checks whether input string are valid utf8 string or not.
// If the string are invalid utf8 string, this method will returns error.
//
func NewLexer(input string, initial flow.Flow) (base.Lexer, error) {

	if !utf8.ValidString(input) {
		return nil, InvalidUtfInput()
	}

	lexer := GenLexer{
		inner:  []rune(input),
		cursor: base.ZeroCursor(),
		flow:   initial,
		tokens: make(chan token.Token, 5),
	}

	return &lexer, nil
}

// DefaultLexer: create new lexer based on string by assuming
//               first state will be a text.
//
// This method checks whether input string are valid utf8 string or not.
// If the string are invalid utf8 string, this method will returns error.
//
func DefaultLexer(input string) (base.Lexer, error) {
	return NewLexer(input, flow.LexTemplate)
}

// Cursor: Get current cursor
//
func (l GenLexer) Cursor() base.Cursor {
	return l.cursor
}

// CursorMut: Get mutable cursor referrence.
//
// Being used for `in-place` update.
//
func (l *GenLexer) CursorMut() *base.Cursor {
	return &l.cursor
}

// HasNext: Check whether there is next state function or not
//
func (l GenLexer) HasNext() bool {

	return l.flow != nil
}

// Next: Get next token in the lexer queue.
//
// This method are already safe since it calls Lexer#HasNext internally.
//
// Will returns error if
//
func (l *GenLexer) Next() (token.Token, error) {
	var err error
	var ok bool
	var t token.Token

	// no next flow
	if !l.HasNext() {
		return token.Token{}, UnavailableFlow()
	}

	for {

		// means no next flow
		if l.flow == nil {
			return token.Token{}, nil
		}

		l.flow, err = l.flow(l)

		if err != nil {
			return token.Token{}, err
		}

		log.Printf("next flow: %v", l.flow)

		select {
		case t, ok = <-l.tokens:
			log.Printf("token from l.tokens: %v\n", t)

			if ok {
				return t, nil
			} else {
				return token.Token{}, fmt.Errorf("channel is closed")
			}
		default:
			log.Printf("default")
			continue
		}

	}
}

// Advance: Advance the cursor in lexer.
//
func (l *GenLexer) Advance() {
	// beware cursor.Next ~ lexer.Advance
	l.cursor.Next()
}

// Available: Check availability of next character
//
func (l GenLexer) Available() bool {
	return l.cursor.Current()+1 < len(l.inner)
}

// Emit : Emit last token being scanned
//
//
func (l GenLexer) LastToken(tt token.Type) token.Token {
	cursor := l.Cursor()

	// need l.current-1, since it's being prefixed with the next token
	value := l.inner[cursor.Start() : l.cursor.Current()-1]

	t := token.Token{tt, value}
	// l.tokens <- token.Token{tt, value}
	return t
}

func (l *GenLexer) Emit(tt token.Type) {
	t := l.LastToken(tt)

	log.Printf("t: %v", t)

	l.tokens <- t

	// assign l.start with l.current, means
	// the range (l.start..l.current) already being
	// consumed by the lexer
	l.CursorMut().Advance()
}

func (l GenLexer) RunesAhead() []rune {
	return l.inner[l.cursor.Current():]
}

func (l GenLexer) CurrentRune() rune {
	return l.inner[l.cursor.Current()]
}

//
// Note: This advance current cursor
//
func (l GenLexer) TakeWhile(conds ...token.RuneCond) {

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
func (l GenLexer) AcceptAll(conds ...token.RuneCond) bool {
	flag := true

	for _, cond := range conds {
		flag = flag && l.Accept(cond)
	}

	return flag
}

//
// Note: This advance current cursor
//
func (l GenLexer) AcceptUntil(conds ...token.RuneCond) bool {
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
func (l GenLexer) Accept(cond token.RuneCond) bool {
	if cond(l.CurrentRune()) {
		if l.Available() {
			l.Advance()
			return true
		}
	}

	// first rune is not an accepted character
	return false
}

func (l GenLexer) Ignore(cond token.RuneCond) {
	for {
		ch := l.CurrentRune()
		if cond(ch) && l.Available() {
			l.Advance()
		} else {
			break
		}
	}
}

func (l GenLexer) String() string {
	return fmt.Sprintf("<lexer{cursor=%s,state=%s}>", l.Cursor(), l.flow)
}

func (l GenLexer) Close() error {
	close(l.tokens)
	log.Println("lexer is closed")
	return nil
}
