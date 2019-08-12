package parser

import (
	"github.com/zerosign/tmpl/base"
	"github.com/zerosign/tmpl/token"
	"github.com/zerosign/tmpl/ast"
)

type Parser struct {
	lex *base.Lexer
	root ast.Node
	tokens []*token.Token
}
