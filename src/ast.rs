// use combine::stream::state::SourcePosition;
use std::collections::HashMap;

//
// Comment repr type.
//
#[derive(Debug, PartialEq)]
pub struct Comment<'a> {
    value: &'a str,
}

//
// Type repr type.
//
#[derive(Debug, PartialEq)]
pub struct Text<'a> {
    value: &'a str,
    // position: (usize, usize),
}

//
// Block sum repr type.
//
// It wraps :
//
// - comment (`{# ... #}`),
// - statement (`{{ ... }}`),
// - expression (`{= ... =}`)
// - `tmpl` text block (other than above)
//
#[derive(Debug, PartialEq)]
pub enum Block<'a> {
    Comment(Comment<'a>),
    Statement(Box<Statement<'a>>),
    Expression(Expression),
    // Text(Text),
}

//
// Expression sum type repr.
//
// Currently, expression sum types only contains :
// - Literal (static value)
// - Functional call value
// - TODO: Boolean expr
//
#[derive(Debug, PartialEq)]
pub enum Expression {
    Value(Literal),
    // FunctionCall(FunctionCall),
    BoolExpr,
}

//
// TODO: how to represents function calls ?
// TODO: need to create struct that wraps rust function at runtime,
//       this also contains compile time types informations. Maybe using macro?
//
// #[derive(Debug)]
// pub struct FunctionCall<F> where F: FnMut {
//     name: String,
//     method: F,
//     args: HashMap<String, Value>,
//     types: HashMap<String, TypeKind>,
// }

//
// Number representations in `tmpl`.
//
#[derive(Debug, PartialEq)]
pub enum Number {
    Integer(i64),
    Double(f64),
}

// Option type repr
type Optional<T> = Option<Box<T>>;

// Most literal types.
//
// - Number ~ f64 and i64 (no unsigned integer)
// - String ~ String (owned type)
// - Bool ~ boolean type
// - Dictionary
// - Array
// - Optional type
//
#[derive(Debug, PartialEq)]
pub enum Literal {
    Number(Number),
    String(String),
    Bool(bool),
    Dictionary(HashMap<String, Literal>),
    Array(Vec<Literal>),
    Optional(Optional<Literal>),
}

#[derive(Debug, PartialEq)]
pub enum TypeKind {
    Bool,
    Number,
    String,
    Dictionary,
    Array,
}

//
// `ident`-like types.
//
#[derive(Debug, PartialEq)]
pub enum Ident {
    Ident(String),
    TypeDecl(String),
    MacroIdent(String),
}

//
// Statement sum types.
//
// It wraps :
//
// - macro statement
// - iterator statement
// - logical statement
//
#[derive(Debug, PartialEq)]
pub enum Statement<'a> {
    MacroStmt,
    IteratorStmt(IteratorStmt<'a>),
    LogicalStmt(LogicalStmt),
}

#[derive(Debug, PartialEq)]
pub struct TypeInfo {
    name: String,
    kind: TypeKind,
    optional: bool,
}

#[derive(Debug, PartialEq)]
pub struct IteratorStmt<'a> {
    key: TypeInfo,
    value: TypeInfo,
    inner: Block<'a>,
}

#[derive(Debug, PartialEq)]
pub enum LogicalStmt {
    // Start
// Repetition
// Edge
}
