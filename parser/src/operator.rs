use combine::{
    char::string, error::ParseError, parser::choice::choice, unexpected_any, value, Parser, Stream,
};

use std::convert::TryFrom;

use crate::{ast, util::lex};

#[inline]
pub fn arithmetic_op<Input>() -> impl Parser<Input, Output = ast::ArithmOp>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    let operator = lex(choice((
        string("+"),
        string("-"),
        string("*"),
        string("/"),
        string("%"),
    )));

    // TODO: use crate::error::ParseError
    operator.then(move |s| match ast::ArithmOp::try_from(s) {
        Ok(v) => value(v).left(),
        _ => unexpected_any(s).right(),
    })
}

#[inline]
pub fn logical_op<Input>() -> impl Parser<Input, Output = ast::LogicalOp>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    let operator = lex(choice((
        string("!="),
        string("=="),
        string(">"),
        string("<"),
        string(">="),
        string("<="),
    )));

    // TODO: use crate::error::ParseError
    operator.then(move |s| match ast::LogicalOp::try_from(s) {
        Ok(v) => value(v).left(),
        _ => unexpected_any(s).right(),
    })
}

#[inline]
pub fn bool_op<Input>() -> impl Parser<Input, Output = ast::BoolOp>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    let operator = lex(choice((string("||"), string("&&"))));

    // TODO: use crate::error::ParseError
    operator.then(move |s| match ast::BoolOp::try_from(s) {
        Ok(v) => value(v).left(),
        _ => unexpected_any(s).right(),
    })
}

#[test]
fn test_arithmetic_op() {
    for op in ast::ArithmOp::all().iter() {
        assert!(arithmetic_op().parse(*op).is_ok());
    }

    assert!(arithmetic_op().parse("+ ").is_ok());
    assert!(arithmetic_op().parse(" + ").is_err());
}

#[test]
fn test_logical_op() {
    for op in ast::LogicalOp::all().iter() {
        assert!(logical_op().parse(*op).is_ok());
    }

    assert!(logical_op().parse("== ").is_ok());
    assert!(logical_op().parse(" == ").is_err());
}

#[test]
fn test_bool_op() {
    for op in ast::BoolOp::all().iter() {
        assert!(bool_op().parse(*op).is_ok());
    }

    assert!(bool_op().parse("|| ").is_ok());
    assert!(bool_op().parse(" || ").is_err());
}
