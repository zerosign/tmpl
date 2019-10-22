use crate::error::ParseError;
use std::{any::TypeId, convert::TryFrom};
use tmpl_value::types::Literal;

#[derive(Debug, Clone, PartialEq)]
pub enum LogicalOp {
    NEQ,
    EQ,
    GT,
    LT,
    GTE,
    LTE,
}

impl LogicalOp {
    #[inline]
    pub fn operator(self) -> &'static str {
        match self {
            Self::NEQ => "!=",
            Self::EQ => "==",
            Self::GT => ">",
            Self::LT => "<",
            Self::GTE => ">=",
            Self::LTE => "<=",
        }
    }

    #[inline]
    pub const fn all() -> [&'static str; 6] {
        ["!=", "==", ">", "<", ">=", "<="]
    }
}

impl<'a> TryFrom<&'a str> for LogicalOp {
    type Error = ParseError<'a>;

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "!=" => Ok(Self::NEQ),
            "==" => Ok(Self::EQ),
            ">" => Ok(Self::GT),
            "<" => Ok(Self::LT),
            ">=" => Ok(Self::GTE),
            "<=" => Ok(Self::LTE),
            _ => Err(ParseError::operator(op, TypeId::of::<Self>())),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum ArithmOp {
    Add,
    Subtract,
    Multiply,
    Divide,
    Modulo,
}

impl ArithmOp {
    #[inline]
    pub fn operator(self) -> &'static str {
        match self {
            Self::Add => "+",
            Self::Subtract => "-",
            Self::Multiply => "*",
            Self::Divide => "/",
            Self::Modulo => "%",
        }
    }

    #[inline]
    pub const fn all() -> [&'static str; 5] {
        ["+", "-", "*", "/", "%"]
    }
}

impl<'a> TryFrom<&'a str> for ArithmOp {
    type Error = ParseError<'a>;

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "+" => Ok(Self::Add),
            "-" => Ok(Self::Subtract),
            "*" => Ok(Self::Multiply),
            "/" => Ok(Self::Divide),
            "%" => Ok(Self::Modulo),
            _ => Err(ParseError::operator(op, TypeId::of::<Self>())),
        }
    }
}

#[derive(Debug, Clone, PartialEq, Copy)]
pub enum BoolOp {
    Or,
    And,
}

impl BoolOp {
    #[inline]
    pub fn operator(self) -> &'static str {
        match self {
            Self::Or => "||",
            Self::And => "&&",
        }
    }

    #[inline]
    pub const fn all() -> [&'static str; 2] {
        ["||", "&&"]
    }
}

impl<'a> TryFrom<&'a str> for BoolOp {
    type Error = ParseError<'a>;

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "||" => Ok(Self::Or),
            "&&" => Ok(Self::And),
            _ => Err(ParseError::operator(op, TypeId::of::<Self>())),
        }
    }
}

//
// Comment repr type.
//
#[derive(Debug, Clone, PartialEq)]
pub struct Comment<'a> {
    value: &'a str,
}

impl<'a> Comment<'a> {
    pub fn new(v: &'a str) -> Comment {
        Comment { value: v }
    }
}

//
// Type repr type.
//
#[derive(Debug, Clone, PartialEq)]
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
#[derive(Debug, Clone, PartialEq)]
pub enum Block<'a> {
    Comment(Comment<'a>),
    Statement(Statement<'a>),
    Expr(Expr),
    Text(Text<'a>),
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
// TODO: how to represents macro calls ?
//
// currently, macro call will be used for including another
// template file into this template file. ( include!(...) ).
//
// pub struct MacroCall<F> where F: FnMut<..?> {
//     name: String,
//     method: F,
//     args: HashMap<_, _>,
//     types: HashMap<_, _>,
//     position: SourcePosition,
// }

#[derive(Debug, Clone, PartialEq)]
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
#[derive(Debug, Clone, PartialEq)]
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
#[derive(Debug, Clone, PartialEq)]
pub enum Statement<'a> {
    MacroStmt,
    IteratorStmt(IteratorStmt<'a>),
    LogicalStmt(LogicalStmt<'a>),
}

#[derive(Debug, Clone, PartialEq)]
pub struct TypeInfo {
    name: String,
    kind: TypeKind,
    optional: bool,
}

#[derive(Debug, Clone, PartialEq)]
pub struct IteratorStmt<'a> {
    key: TypeInfo,
    value: TypeInfo,
    inner: Box<Block<'a>>,
}

#[derive(Debug, Clone, PartialEq)]
pub struct LogicalClause<'a> {
    inner: Vec<BoolExpr>,
    block: Option<Box<Block<'a>>>,
}

//
// Note: Logical statement should have alternative Else clause.
//
//
#[derive(Debug, Clone, PartialEq)]
pub enum LogicalStmt<'a> {
    Base(LogicalClause<'a>),
    Alt(Option<Box<Block<'a>>>),
}

#[derive(Debug, Clone, PartialEq)]
pub enum Expr {
    BoolExpr(BoolExpr),
    ArithmExpr(ArithmExpr),
    LogicalExpr(LogicalExpr),
    NegatedExpr(NegatedExpr),
    NegativeExpr(ArithmExpr),
    LiteralExpr(Literal),
}

#[derive(Debug, Clone, PartialEq)]
pub struct NegatedExpr {
    v: Box<BoolExpr>,
}

#[derive(Debug, Clone, PartialEq)]
pub struct BoolExpr {
    lhs: Box<LogicalExpr>,
    rhs: Box<LogicalExpr>,
    op: BoolOp,
}

#[derive(Debug, Clone, PartialEq)]
pub struct ArithmExpr {
    lhs: Box<Expr>,
    rhs: Box<Expr>,
    op: ArithmOp,
}

impl TryFrom<((Literal, ArithmOp), Literal)> for ArithmExpr {
    type Error = ();

    #[inline]
    fn try_from(v: ((Literal, ArithmOp), Literal)) -> Result<ArithmExpr, Self::Error> {
        match v {
            ((Literal::Number(l), o), Literal::Number(r)) => Ok(ArithmExpr {
                lhs: Box::new(Expr::LiteralExpr(Literal::Number(l))),
                rhs: Box::new(Expr::LiteralExpr(Literal::Number(r))),
                op: o,
            }),
            _ => Err(()),
        }
    }
}

impl TryFrom<((Literal, ArithmOp), ArithmExpr)> for ArithmExpr {
    type Error = ();

    #[inline]
    fn try_from(v: ((Literal, ArithmOp), ArithmExpr)) -> Result<ArithmExpr, Self::Error> {
        match v {
            ((Literal::Number(l), o), r) => Ok(ArithmExpr {
                lhs: Box::new(Expr::LiteralExpr(Literal::Number(l))),
                rhs: Box::new(Expr::ArithmExpr(r)),
                op: o,
            }),
            _ => Err(()),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub struct LogicalExpr {
    lhs: Box<Expr>,
    rhs: Box<Expr>,
    op: LogicalOp,
}
