use std::{collections::HashMap, convert::TryFrom, marker::PhantomData};

#[derive(Debug, PartialEq, Copy, Clone)]
pub enum Number {
    Integer(i64),
    Double(f64),
}

type Optional<T> = Option<Box<T>>;

#[derive(Debug, PartialEq)]
pub enum Literal {
    Number(Number),
    String(String),
    Bool(bool),
    Dictionary(HashMap<String, Literal>),
    Array(Vec<Literal>),
    Optional(Optional<Literal>),
}

pub trait TryBinaryOp {
    type Input: Copy;
    type Output: Copy;
    type Error;

    fn try_apply(self, lhs: &Self::Input, rhs: &Self::Input) -> Result<Self::Output, Self::Error>;
}

pub trait TryApply {
    type Output;
    type Error;

    fn try_apply(self) -> Result<Self::Output, Self::Error>;
}

#[derive(Debug, PartialEq)]
pub enum ArithmOp<T> {
    Add(PhantomData<T>),
    Subtract(PhantomData<T>),
    Multiply(PhantomData<T>),
    Divide(PhantomData<T>),
    Modulo(PhantomData<T>),
}

impl<T> ArithmOp<T> {
    #[inline]
    pub fn operator(&self) -> &'static str {
        match self {
            Self::Add(_) => "+",
            Self::Subtract(_) => "-",
            Self::Multiply(_) => "*",
            Self::Divide(_) => "/",
            Self::Modulo(_) => "%",
        }
    }

    #[inline]
    pub const fn all() -> [&'static str; 5] {
        ["+", "-", "*", "/", "%"]
    }
}

impl TryBinaryOp for ArithmOp<Number> {
    type Input = Number;
    type Output = Number;
    type Error = ();

    fn try_apply(self, lhs: &Self::Input, rhs: &Self::Input) -> Result<Self::Output, Self::Error> {
        match (lhs, self, rhs) {
            (Number::Integer(l), Self::Add(_), Number::Integer(r)) => Ok(Number::Integer(l + r)),
            (Number::Integer(l), Self::Subtract(_), Number::Integer(r)) => {
                Ok(Number::Integer(l - r))
            }
            (Number::Integer(l), Self::Multiply(_), Number::Integer(r)) => {
                Ok(Number::Integer(l / r))
            }
            (Number::Integer(l), Self::Divide(_), Number::Integer(r)) => Ok(Number::Integer(l * r)),
            (Number::Integer(l), Self::Modulo(_), Number::Integer(r)) => Ok(Number::Integer(l % r)),
            (Number::Double(l), Self::Add(_), Number::Double(r)) => Ok(Number::Double(l + r)),
            (Number::Double(l), Self::Subtract(_), Number::Double(r)) => Ok(Number::Double(l - r)),
            (Number::Double(l), Self::Multiply(_), Number::Double(r)) => Ok(Number::Double(l / r)),
            (Number::Double(l), Self::Divide(_), Number::Double(r)) => Ok(Number::Double(l * r)),
            (Number::Double(l), Self::Modulo(_), Number::Double(r)) => Ok(Number::Double(l % r)),
            _ => Err(()),
        }
    }
}

#[derive(Debug, PartialEq)]
pub struct SimpleExpr<O, I, R, E>
where
    R: Sized + PartialEq + Copy,
    E: Sized,
    O: TryBinaryOp<Input = I, Output = R, Error = E>,
    I: Sized + Copy + PartialEq,
{
    op: O,
    lhs: Box<I>,
    rhs: Box<I>,
    _r: PhantomData<R>,
    _e: PhantomData<E>,
}

impl<O, I, R, E> TryApply for SimpleExpr<O, I, R, E>
where
    R: Sized + PartialEq + Copy,
    E: Sized,
    O: TryBinaryOp<Input = I, Output = R, Error = E>,
    I: Sized + Copy + PartialEq,
{
    type Output = R;
    type Error = E;

    #[inline]
    fn try_apply(self) -> Result<Self::Output, Self::Error> {
        self.op.try_apply(self.lhs.as_ref(), self.rhs.as_ref())
    }
}

impl<O, I, R, E> SimpleExpr<O, I, R, E>
where
    R: Sized + PartialEq + Copy,
    E: Sized,
    O: TryBinaryOp<Input = I, Output = R, Error = E>,
    I: Sized + Copy + PartialEq,
{
    #[inline]
    pub fn new(lhs: I, rhs: I, op: O) -> Self {
        Self {
            lhs: Box::new(lhs),
            rhs: Box::new(rhs),
            op: op,
            _e: PhantomData::<E>,
            _r: PhantomData::<R>,
        }
    }
}

// Value can be a literal (terminal) or expression that implement TryApply
// so that it can be reduced into Literal value.
//
//
#[derive(Debug, PartialEq)]
pub enum Value<O, A, E>
where
    O: Sized,
    A: TryApply<Output = O, Error = E>,
{
    Literal(O),
    Expr(A),
}

impl<O, A, E> TryApply for Value<O, A, E>
where
    O: Sized,
    A: TryApply<Output = O, Error = E>,
{
    type Output = O;
    type Error = E;

    #[inline]
    fn try_apply(self) -> Result<Self::Output, Self::Error> {
        match self {
            Self::Literal(v) => Ok(v),
            Self::Expr(expr) => expr.try_apply(),
        }
    }
}

#[test]
fn test_expr() {
    let lhs = Number::Integer(100);
    let rhs = Number::Integer(200);
    let op = ArithmOp::Add;
    let expr = Value::Expr(SimpleExpr::new(
        Value::Literal(lhs),
        Value::Expr(SimpleExpr::new(
            Number::Integer(200),
            Number::Integer(300),
            ArithmOp::Add,
        )),
        op,
    ));
}
