use combine::{
    char::{char, spaces, string},
    error::{Consumed, ParseError, StreamError},
    parser::{
        choice::choice,
        error::unexpected,
        item::{position, value},
    },
    Parser, Stream,
};

use std::convert::TryFrom;

use crate::{ast, literal};

#[inline]
pub fn arithmetic_op<I>() -> impl Parser<Input = I, Output = ast::ArithmOp>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    let operator = choice((
        string("+"),
        string("-"),
        string("*"),
        string("/"),
        string("%"),
    ));

    // let operators = ast::ArithmOp::all()
    //     .iter()
    //     .map(|s| string(*s))
    //     .collect::<Vec<_>>();

    // let operator = choice(operators.into());

    // TODO: please convert error rather doing an unwrap like this
    operator.map(move |s| ast::ArithmOp::try_from(s).unwrap())
}

#[inline]
pub fn logical_op<I>() -> impl Parser<Input = I, Output = ast::LogicalOp>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    let operator = choice((
        string("!="),
        string("=="),
        string(">"),
        string("<"),
        string(">="),
        string("<="),
    ));

    // TODO: please convert error rather doing an unwrap like this
    operator.map(move |s| ast::LogicalOp::try_from(s).unwrap())
}

#[inline]
pub fn bool_op<I>() -> impl Parser<Input = I, Output = ast::BoolOp>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    let operator = choice((string("||"), string("&&")));

    // TODO: please convert error rather doing an unwrap like this
    operator.map(move |s| ast::BoolOp::try_from(s).unwrap())
}
