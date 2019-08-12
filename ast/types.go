package ast

// Location: location information of the node
//
//
type Location struct {
	Position, Line int
}

// NodeType: node type of node
//
type NodeType int

const (
// TODO: @zerosign set all nodetypes in here
)

// Node: interface of the node
//
type Node interface {
	// node type
	Type() NodeType

	// location of node in original input string
	Location() Location

	// string representation, used for debugging
	String() string

	// accepts a visitor
	Accept(Visitor) (interface{}, []error)
}

type CommentStatement struct {
	NodeType
	Location
	Value string
}

type StringLiteral struct {
	NodeType
	Location
	Value string
}

type BooleanLiteral struct {
	NodeType
	Location
	Value    bool
	Original string
}

type NumberLiteral struct {
	NodeType
	Location
	DoubleValue  float64
	IntegerValue int64
	Original     string
}

type Statement interface{}

type LoopStmt struct {
	children []Statement
}

type Visitor interface {

	// visit comment
	VisitComment(CommentStatement) (interface{}, []error)

	// visit string literal
	VisitString(StringLiteral) (interface{}, []error)

	// visit number literal
	VisitNumber(NumberLiteral) (interface{}, []error)

	// visit boolean literal
	VisitBoolean(BooleanLiteral) (interface{}, []error)
}
