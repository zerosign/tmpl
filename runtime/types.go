package runtime

import (
	// hasher "github.com/cnf/structhash"
	"github.com/zerosign/tmpl/ast"
	"github.com/zerosign/tmpl/function"
)


const (
	FieldTemplatePath = "template_path"
	FieldTemplateName = "template_name"
)


type Template struct {
	ctx  function.Context
	root ast.Node
	hash string
}

func (t Template) TemplateName() string {
	return t.ctx[FieldTemplateName].(string)
}

func (t Template) TemplatePath() string {
	return t.ctx[FieldTemplatePath].(string)
}
