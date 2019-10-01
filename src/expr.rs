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
    literal, operator, util, value,
};
use combine::{between, char::char, choice, ParseError, Parser, Stream};
use std::convert::TryFrom;

const BlockClause: &'static (char, char) = &('(', ')');

pub fn block<I, B, O>(b: B, clause: (char, char)) -> impl Parser<Input = I, Output = O>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
    B: Parser<Input = I, Output = O>,
{
    between(char(clause.0), char(clause.1), b)
}

fn simple_arith_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    util::lex::<_, _, Literal>(literal::number_literal())
        .and(util::lex::<_, _, ArithmOp>(operator::arithmetic_op()))
        .and(util::lex::<_, _, Literal>(literal::number_literal()))
        // TODO: need to use and_then
        .map(|x| ArithmExpr::try_from(x).unwrap())
}

fn expand_arith_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    util::lex::<_, _, Literal>(literal::number_literal())
        .and(util::lex::<_, _, ArithmOp>(operator::arithmetic_op()))
        .and(util::lex::<_, _, ArithmExpr>(simple_arith_expr()))
        // TODO: need to use and_then
        .map(|x| ArithmExpr::try_from(x).unwrap())
}

//
// arithm_expr ->
//    block ( arithm_expr )
//
//
// #[inline]
// pub fn arithmetic_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     let x = simple_arith_expr();
//     let y = expand_arith_expr();

//     choice([x, y])
// }

#[test]
fn test_simple_arith_expr() {
    println!("{:?}", simple_arith_expr().parse("1 + 2 "));
}
