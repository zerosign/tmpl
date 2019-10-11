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
    ast::{ArithmExpr, ArithmOp, Expr, Literal},
    literal, operator,
    util::{lex, para},
};
use combine::{choice, unexpected_any, value, ParseError, Parser, Stream};
use std::convert::TryFrom;
// use std::convert::TryFrom;

// arithmetic
//
// lhs op rhs
// lhs ~ arithmetic_expr
// rhs ~ arithmetic_expr
//
// 3 choices
// number_literal op number_literal
// number_literal op arithmetic_expr
// arithmetic_expr op number_literal

fn arithmetic_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    choice([literal::number_literal()
        .and(operator::arithmetic_op())
        .and(literal::number_literal())
        .then(|x| match ArithmExpr::try_from(x) {
            Ok(v) => value(v).left(),
            _ => unexpected_any('x').right(),
        })])
}

#[test]
fn test_simple_arith_expr() {
    let tests = vec!["1 + 2", "1 + (2)", "((1)) + 2"];
    let expected =
        ArithmExpr::try_from(((Literal::integer(1), ArithmOp::Add), Literal::integer(2))).unwrap();

    for test in tests {
        assert!(
            arithmetic_expr().parse(test).is_ok(),
            // Ok((expected.clone(), "")),
            format!("error in {:?}", test)
        );
    }

    // assert_eq!(r, Ok((expected, "")));
}
