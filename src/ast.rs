// use combine::stream::state::SourcePosition;

use crate::error::ParseError;
use std::{any::TypeId, collections::HashMap, convert::TryFrom, iter::IntoIterator};

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
    pub fn operator(&self) -> &'static str {
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
//
// Number representations in `tmpl`.
//
#[derive(Debug, Clone, PartialEq)]
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
// - Optional type
//
#[derive(Debug, Clone, PartialEq)]
pub enum Literal {
    Number(Number),
    String(String),
    Bool(bool),
    Optional(Optional<Literal>),
}

impl Literal {
    #[inline]
    pub(crate) fn integer<V>(v: V) -> Literal
    where
        V: Into<i64>,
    {
        Literal::Number(Number::Integer(v.into()))
    }

    #[inline]
    pub(crate) fn double<V>(v: V) -> Literal
    where
        V: Into<f64>,
    {
        Literal::Number(Number::Double(v.into()))
    }

    #[inline]
    pub(crate) fn string<V>(v: V) -> Literal
    where
        V: Into<String>,
    {
        Literal::String(v.into())
    }

    #[inline]
    pub(crate) fn bool<V>(v: V) -> Literal
    where
        V: Into<bool>,
    {
        Literal::Bool(v.into())
    }
}

macro_rules! literal_conv {
    ($($conv:path => [$($src:ty),*]),*) => {
        $($(impl From<$src> for Literal {

            #[inline]
            fn from(v: $src) -> Self {
                $conv(v)
            }
        })*)*
    }
}

//
// Note: we don't support u64 at this point.
//
literal_conv!(
    Literal::integer => [u8, u16, u32, i8, i16, i32, i64],
    Literal::double  => [f32, f64],
    Literal::string  => [String, &'static str],
    Literal::bool    => [bool]
);

// value =
//   value = literal
//   dictionary<string, value>
//   array<value>
//
// - Dictionary
// - Array
//
#[derive(Debug, PartialEq)]
pub enum Value {
    Literal(Literal),
    Dictionary(HashMap<String, Value>),
    Array(Vec<Value>),
}

impl Value {
    pub fn literal_array<I, V>(i: I) -> Value
    where
        I: IntoIterator<Item = V>,
        V: Into<Literal>,
    {
        Value::Array(
            i.into_iter()
                .map(|v| Value::Literal(V::into(v)))
                .collect::<Vec<Value>>(),
        )
    }
}

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
    NegatifExpr(ArithmExpr),
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
