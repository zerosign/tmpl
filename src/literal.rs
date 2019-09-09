use crate::ast::{Literal, Number};
use combine::{
    char::{char, digit, spaces, string},
    error::{Consumed, ParseError},
    parser,
    parser::{
        choice::{choice, optional},
        error::unexpected_any,
        function::parser as fparser,
        item::{any, satisfy_map, value},
        repeat::{many, many1, sep_by},
        sequence::between,
    },
    stream::Stream,
    Parser,
};

#[inline]
pub fn bool<I>() -> impl Parser<Input = I, Output = Literal>
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
pub fn quoted_str_<I>() -> impl Parser<Input = I, Output = String>
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
pub fn quoted_str<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    quoted_str_().map(|s| Literal::String(s))
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
        if s.len() > 1 && s.starts_with("0") {
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
fn fractional<'a, I>() -> impl Parser<Input = I, Output = f64>
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
pub fn number<'a, I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    optional(char('-'))
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

#[inline]
pub fn dict<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    let field = (quoted_str_(), char(':').skip(spaces()), static_value()).map(|t| (t.0, t.2));
    let fields = sep_by(field, char(',').skip(spaces()));

    between(char('{').skip(spaces()), char('}').skip(spaces()), fields)
        .map(Literal::Dictionary)
        .expected("dictionary")
}

// pub fn optional<T>(p: P) -> impl Parser<Input = I, Output = Literal::Optional>
// where
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     string("null")
// }

parser! {
    #[inline]
    pub fn static_value_[I]()(I) -> Literal
    where [ I: Stream<Item = char> ]
    {
        let array = between(
            char('[').skip(spaces()),
            char(']').skip(spaces()),
            sep_by(static_value(), char(',').skip(spaces()))
        ).map(Literal::Array);

        choice((bool(), quoted_str(), array, dict(), number()))
    }
}

#[inline]
pub fn static_value<I>() -> impl Parser<Input = I, Output = Literal>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    static_value_()
}
