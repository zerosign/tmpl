//!
//! ```
//! expr ->
//!   logical_expr = == != >= <= < > (Eq, Order, expr)
//!   bool_expr = &&  || (can be used for bool, logical_expr, bool_expr)
//!   arith_expr = + - / * % (can be used for number)
//!   negated_expr = ! (bool_expr)
//!   negative_expr = - (arith_expr)
//! ```
//!
//! block -> '(' expr ')'
//!
use crate::{
    ast::{ArithmExpr, ArithmOp, Literal},
    literal, operator,
    util::{lex, para},
};
use combine::{ParseError, Parser, Stream};
use std::convert::TryFrom;

fn simple_arith_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    para(literal::number_literal())
        .and(para(operator::arithmetic_op()))
        .and(para(literal::number_literal()))
        // TODO: need to use and_then
        .map(|x| ArithmExpr::try_from(x).unwrap())
}

fn expand_arith_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    para(literal::number_literal())
        .and(lex::<_, _, ArithmOp>(operator::arithmetic_op()))
        .and(para(simple_arith_expr()))
        // TODO: need to use and_then
        .map(|x| ArithmExpr::try_from(x).unwrap())
}

// arithm_expr ->
//    block ( arithm_expr )
//
// #[inline]
// pub fn arithmetic_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
// }

#[test]
fn test_simple_arith_expr() {
    // println!("{:?}", simple_arith_expr().parse("1 + 2"));

    let data = vec!["1 + 2", "1 + (2)", "(((1)) + 2)", "1 + 2 + (2 + 1 + (-1))"];
}
