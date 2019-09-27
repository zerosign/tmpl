// use combine::stream::state::SourcePosition;
use crate::{
    error::CastError,
    types::{Apply, BinaryOp},
};

use std::{
    any::TypeId,
    collections::HashMap,
    convert::{TryFrom, TryInto},
};

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
    pub fn operator(&self) -> &'static str {
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
    type Error = ();

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "!=" => Ok(Self::NEQ),
            "==" => Ok(Self::EQ),
            ">" => Ok(Self::GT),
            "<" => Ok(Self::LT),
            ">=" => Ok(Self::GTE),
            "<=" => Ok(Self::LTE),
            _ => Err(()),
        }
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum ArithmOp {
    Add,
    Substract,
    Multiply,
    Divide,
    Modulo,
}

impl ArithmOp {
    #[inline]
    pub fn operator(&self) -> &'static str {
        match self {
            Self::Add => "+",
            Self::Substract => "-",
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
    type Error = ();

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "+" => Ok(Self::Add),
            "-" => Ok(Self::Substract),
            "*" => Ok(Self::Multiply),
            "/" => Ok(Self::Divide),
            "%" => Ok(Self::Modulo),
            _ => Err(()),
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
    pub fn operator(&self) -> &'static str {
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
    type Error = ();

    #[inline]
    fn try_from(op: &'a str) -> Result<Self, Self::Error> {
        match op {
            "||" => Ok(Self::Or),
            "&&" => Ok(Self::And),
            _ => Err(()),
        }
    }
}

impl<T> BinaryOp<T, bool> for BoolOp
where
    T: AsRef<Boolean>,
{
    fn apply(&self, lhs: T, rhs: T) -> bool {
        match (lhs.as_ref(), self, rhs.as_ref()) {
            (Boolean::Literal(ref l), Self::Or, Boolean::Literal(ref r)) => *l || *r,
            (Boolean::Literal(ref l), Self::Or, Boolean::Expr(ref r)) => {
                *l || r.op.apply(r.lhs, r.rhs)
            }
            (Boolean::Expr(l), Self::Or, Boolean::Literal(r)) => l.op.apply(l.lhs, l.rhs) || *r,
            (Boolean::Expr(l), Self::Or, Boolean::Expr(r)) => {
                l.op.apply(l.lhs, l.rhs) || r.op.apply(r.lhs, r.rhs)
            }
            (Boolean::Literal(l), Self::And, Boolean::Literal(r)) => *l && *r,
            (Boolean::Literal(l), Self::And, Boolean::Expr(r)) => *l && r.op.apply(r.lhs, r.rhs),
            (Boolean::Expr(l), Self::And, Boolean::Literal(r)) => l.op.apply(l.lhs, l.rhs) && *r,
            (Boolean::Expr(l), Self::And, Boolean::Expr(r)) => {
                l.op.apply(l.lhs, l.rhs) && r.op.apply(r.lhs, r.rhs)
            }
        }
    }
}

impl BinaryOp<Box<Boolean>, Box<Boolean>> for BoolOp {
    fn apply(&self, lhs: Box<Boolean>, rhs: Box<Boolean>) -> Box<Boolean> {
        match (*lhs, self, *rhs) {
            (Boolean::Literal(l), Self::Or, Boolean::Literal(r)) => {
                Box::new(Boolean::Literal(l || r))
            }
            (Boolean::Literal(l), Self::Or, Boolean::Expr(r)) => {
                Box::new(Boolean::Literal(l || r.op.apply(r.lhs, r.rhs)))
            }
            (Boolean::Expr(l), Self::Or, Boolean::Literal(r)) => {
                Box::new(Boolean::Literal(l.op.apply(l.lhs, l.rhs) || r))
            }

            (Boolean::Expr(l), Self::Or, Boolean::Expr(r)) => Box::new(Boolean::Literal(
                l.op.apply(l.lhs, l.rhs) || r.op.apply(r.lhs, r.rhs),
            )),
            (Boolean::Literal(l), Self::And, Boolean::Literal(r)) => {
                Box::new(Boolean::Literal(l && r))
            }
            (Boolean::Literal(l), Self::And, Boolean::Expr(r)) => {
                Box::new(Boolean::Literal(l && r.op.apply(r.lhs, r.rhs)))
            }
            (Boolean::Expr(l), Self::And, Boolean::Literal(r)) => {
                Box::new(Boolean::Literal(l.op.apply(l.lhs, l.rhs) && r))
            }
            (Boolean::Expr(l), Self::And, Boolean::Expr(r)) => Box::new(Boolean::Literal(
                l.op.apply(l.lhs, l.rhs) && r.op.apply(r.lhs, r.rhs),
            )),
        }
    }
}

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
    Text(Text<'a>),
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
    Literal(Literal),
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

impl<'a> TryInto<Boolean> for Literal {
    type Error = CastError<Self>;

    #[inline]
    fn try_into(self) -> Result<Boolean, Self::Error> {
        match self {
            Self::Bool(v) => Ok(Boolean::Literal(v)),
            _ => Err(CastError::IncompatibleCast(self, TypeId::of::<Boolean>())),
        }
    }
}

#[derive(Debug, PartialEq, Copy)]
pub enum Boolean {
    Literal(bool),
    Expr(SimpleExpr<BoolOp, Box<Boolean>>),
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

//
// Note: Logical statement should have alternative Else clause.
//
//
#[derive(Debug, PartialEq)]
pub enum LogicalStmt {
    IfClause(),
    IfElseClause,
    ElseClause(),
}

#[derive(Debug, PartialEq, Copy)]
pub struct SimpleExpr<O, I>
where
    O: BinaryOp<I, I>,
    I: Sized + PartialEq + Copy,
{
    op: O,
    lhs: I,
    rhs: I,
}

impl<O, I> Apply<I> for SimpleExpr<O, I>
where
    O: BinaryOp<I, I>,
    I: Sized + PartialEq + Copy,
{
    fn apply(&self) -> I {
        self.op.apply(self.lhs, self.rhs)
    }
}

// expression ast are recursively defined (it's quite dangerous to be defined).
//
// pub type LogicalExpr = SimpleExpr<Literal, LogicalOp>;
