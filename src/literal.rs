//!
//! Literal parser.
//!
//! ```
//! literal =
//!   number = integer double
//!   string
//!   bool
//! ```
//!
use crate::{
    ast::{Literal, Number},
    util::{lex, para},
};
use combine::{
    char::{char, digit, spaces, string},
    error::{Consumed, ParseError},
    parser::{
        choice::optional,
        error::unexpected_any,
        function::parser as fparser,
        item::{any, satisfy_map, value},
        repeat::{many, many1},
        sequence::between,
    },
    stream::Stream,
    Parser,
};

#[inline]
pub fn bool_literal<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    string("true")
        .map(|_| Literal::Bool(true))
        .or(string("false").map(|_| Literal::Bool(false)))
}

#[inline]
fn escaped_str_chars(c: char) -> Option<char> {
    match c {
        '"' => Some('"'),
        '\\' => Some('\\'),
        '/' => Some('/'),
        'b' => Some('\u{0008}'),
        'f' => Some('\u{000c}'),
        'n' => Some('\n'),
        'r' => Some('\r'),
        't' => Some('\t'),
        _ => None,
    }
}

//
// This was taken from `combine` codes in `benches/json.rs`.
//
//
#[inline]
pub fn escaped_char<I, M>(matcher: M) -> impl Parser<Input = I, Output = char>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
    M: FnMut(I::Item) -> Option<char> + Copy,
{
    // this equal to pointer lookahead but done lazily
    fparser(move |input: &mut I| {
        // scan lazily first char
        let (c, consumed) = any().parse_lazy(input).into_result()?;

        // check whether c ~ '\\'
        // if yes continue matching the next char based on matcher function
        // then return the mapping based on matcher function
        match c {
            '\\' => {
                consumed.combine(move |_| satisfy_map(matcher).parse_stream(input).into_result())
            }
            '"' => Err(Consumed::Empty(I::Error::empty(input.position()).into())),
            _ => Ok((c, consumed)),
        }
    })
}

#[inline]
pub fn raw_string<I>() -> impl Parser<Input = I, Output = String>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    between(
        char('"'),
        char('"').skip(spaces()),
        many(escaped_char(escaped_str_chars)),
    )
    .expected("string")
}

#[inline]
pub fn string_literal<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    raw_string().map(Literal::String)
}

//
//
//
#[inline]
fn integer<I>() -> impl Parser<Input = I, Output = i64>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    many1(digit()).then(|s: String| {
        if s.len() > 1 && s.starts_with('0') {
            unexpected_any('0')
                .message("no `0` before any digit at beginning in integer")
                .right()
        } else if s.starts_with("00") {
            unexpected_any('0')
                .message("no consecutive `0` is allowed at beginning")
                .right()
        } else {
            let int_val = s
                .chars()
                .fold(0, |r, ch| (r * 10) + (ch as i64 - '0' as i64));

            value(int_val).left()
        }
    })
}

//
//
//
#[inline]
fn fractional<I>() -> impl Parser<Input = I, Output = f64>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    many(digit()).then(move |mut digits: String| {
        digits.insert(0, '.');
        digits.insert(0, '0');
        match digits.parse::<f64>() {
            Ok(v) => value(v).left(),
            // TODO(@zerosign): need to know how pass normal error to ParseError
            Err(_) => unexpected_any(' ').right(),
        }
    })
}

//
//
//
#[inline]
pub fn number_literal<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    optional(lex::<_, _, char>(char('-')))
        .map(|c| if c.is_some() { -1 } else { 1 })
        .and(integer())
        .and(optional(char('.').with(fractional())))
        .map(|arg| match arg {
            ((mult, exp), Some(frac)) => {
                Literal::Number(Number::Double((exp as f64 + frac) * (mult as f64)))
            }
            ((mult, exp), None) => Literal::Number(Number::Integer(exp * mult)),
        })
}

#[test]
fn test_number_literal() {
    assert_eq!(
        number_literal().parse("1000000000"),
        Ok((Literal::integer(1000000000), ""))
    );

    assert_eq!(
        number_literal().parse("1000000000.0"),
        Ok((Literal::double(1000000000.0), ""))
    );

    assert_eq!(
        number_literal().parse("-1000000000"),
        Ok((Literal::integer(-1000000000), ""))
    );

    assert_eq!(
        number_literal().parse("-1000000000.0"),
        Ok((Literal::double(-1000000000.0), ""))
    );

    assert_eq!(
        number_literal().parse("-0.0"),
        Ok((Literal::double(-0.0), ""))
    );

    assert_eq!(
        number_literal().parse("0.0"),
        Ok((Literal::double(0.0), ""))
    );

    assert_eq!(
        number_literal().parse("0.011111111"),
        Ok((Literal::double(0.011111111), ""))
    );
}

#[test]
fn test_bool_literal() {
    assert_eq!(bool_literal().parse("true"), Ok((Literal::bool(true), "")));
    assert_eq!(
        bool_literal().parse("false"),
        Ok((Literal::bool(false), ""))
    );
}

#[test]
fn test_string_literal() {
    assert_eq!(
        string_literal().parse(r#""Hello world\ntest!""#),
        Ok((Literal::string("Hello world\ntest!"), ""))
    );

    assert_eq!(
        string_literal().parse(r#""""#),
        Ok((Literal::string(""), ""))
    );

    // TODO: quoted string test
    // assert_eq!(
    //     string_literal().parse("\\\"\\\""),
    //     Ok((Literal::string("\"\""), ""))
    // );

    // TODO: escaped string test
}
