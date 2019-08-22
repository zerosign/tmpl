package parser

import (
	"github.com/zerosign/tmpl/ast"
	"github.com/zerosign/tmpl/lexer"
	"github.com/zerosign/tmpl/token"
)

type Parser struct {
	lex    *lexer.Lexer
	root   ast.Node
	tokens []*token.Token
}
