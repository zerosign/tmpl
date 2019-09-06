#![feature(try_trait)]
extern crate combine;
extern crate env_logger;
extern crate log;

use combine::{
    error::{Consumed, ParseError},
    parser::{
        char::{char, digit, letter, spaces, string},
        choice::{choice, optional},
        error::unexpected_any,
        function::parser,
        item::{any, satisfy_map, value},
        repeat::{many, many1},
        sequence::between,
    },
    stream::state::State,
    Parser, Stream, StreamOnce,
};
use std::ops::Try;

#[derive(Debug, PartialEq)]
pub enum Number {
    Float(f64),
    Integer(i64),
}

#[derive(Debug, PartialEq)]
pub enum Value {
    String(String),
    Bool(bool),
    Number(Number),
    Array(Vec<Box<Value>>),
    Pair { key: String, value: Box<Value> },
}

// #[inline]
// fn quoted_char<I>() -> impl Parser<Input = I, Output = char>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     // escaped_char *>
// }

#[inline]
fn quoted_str<I>() -> impl Parser<Input = I, Output = Value>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    between(char('"'), char('"'), many(letter())).map(|value| Value::str(value))
}

#[inline]
fn integer<I>() -> impl Parser<Input, Output = Value>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    many1(digit()).map(|s: String| {
        let mut value = 0;
        for c in s.chars() {
            value = value * 10 + (c as i64 - '0' as i64);
        }
        Value::Number(Number::Integer(value))
    })
}

#[inline]
fn number<I>() -> impl Parser<Input = I, Output = Value>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    // digit without zero or zero
    let digit_without_zero = digit().then(|ch| {
        if ch != '0' {
            value(ch).left()
        } else {
            unexpected_any(ch)
                .message("'0' shouldn't be allowed")
                .right()
        }
    });

    let non_zero_integer = digit_without_zero().and(many(digit())).map(|(f, rest)| {
        let other = rest.parse::<i64>();
    });

    optional(char('-')).and(integer)
}

#[inline]
fn bool<I>() -> impl Parser<Input = I, Output = Value>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    string("true")
        .map(|_| Value::Bool(true))
        .or(string("false").map(|_| Value::Bool(false)))
}

#[cfg(test)]
mod tests {

    use crate::{bool, combine::Parser, quoted_str, Value};

    #[test]
    fn parse_quoted_str_test() {}

    #[test]
    fn parse_bool_test() {
        let (value, left) = bool().parse("true").unwrap();
        assert_eq!(value, Value::Bool(true));

        let (value, left) = bool().parse("false").unwrap();
        assert_eq!(value, Value::Bool(false));
    }

    #[test]
    fn parse_integer_test() {}
}
