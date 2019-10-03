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
use combine::{choice, ParseError, Parser, Stream};
// use std::convert::TryFrom;

// fn simple_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     literal::number_literal()
//         .and(operator::arithmetic_op())
//         .and(literal::number_literal())
//         .map(|x| ArithmExpr::try_from(x).unwrap())
// }

// fn simple_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr> {
//     literal::number_literal().and(operator::arithmetic_op()).and
// }

// fn expand_arith_expr<I>() -> impl Parser<Input = I, Output = ArithmExpr>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     literal::number_literal()
//         .and(operator::arithmetic_op())
//         .and(simple_arith_expr())
//         // TODO: need to use and_then
//         .map(|x| ArithmExpr::try_from(x).unwrap())
// }

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

    // let tests = vec!["1 + 2", "1 + (2)", "(1 + 2)", "(1) + 2"];
    // let expected =
    //     ArithmExpr::try_from(((Literal::integer(1), ArithmOp::Add), Literal::integer(2))).unwrap();

    // for test in tests {
    //     assert!(
    //         simple_expr().parse(test).is_ok(),
    //         // Ok((expected.clone(), "")),
    //         format!("error in {:?}", test)
    //     );
    // }

    // assert_eq!(r, Ok((expected, "")));
}
