//!
//! ```
//! value =
//!   value = literal
//!   dictionary<string, value>
//!   array<value>
//! ```
//!
use crate::{
    ast::{Literal, Value},
    literal,
    util::para,
};
use combine::{
    char::{char, spaces},
    error::ParseError,
    parser,
    parser::{choice::choice, repeat::sep_by, sequence::between},
    stream::Stream,
    Parser,
};

pub fn literal<Input>() -> impl Parser<Input, Output = Value>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    para::<_, _, Literal>(choice((
        literal::bool_literal(),
        literal::string_literal(),
        literal::number_literal(),
    )))
    .map(|l| Value::Literal(l))
}

#[inline]
pub fn dict<Input>() -> impl Parser<Input, Output = Value>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    let field = (
        literal::raw_string(),
        char(':').skip(spaces()),
        static_value(),
    )
        .map(|t| (t.0, t.2));
    let fields = sep_by(field, char(',').skip(spaces()));

    between(char('{').skip(spaces()), char('}').skip(spaces()), fields)
        .map(Value::Dictionary)
        .expected("dictionary")
}

//
// TODO: optional type
//
// pub fn optional<T>(p: P) -> impl Parser<Input = I, Output = Literal::Optional>
// where
//     I: Stream<Token = char>,
//     I::Error: ParseError<I::Token, I::Range, I::Position>,
// {
//     string("null")
// }

pub fn array<Input>() -> impl Parser<Input, Output = Value>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    between(
        char('[').skip(spaces()),
        char(']').skip(spaces()),
        sep_by(static_value(), char(',').skip(spaces())),
    )
    .map(Value::Array)
}

parser! {
    #[inline]
    pub fn static_value_[Input]()(Input) -> Value
    where [ Input: Stream<Token = char> ]
    {

        choice((literal(), array(), dict()))
    }
}

#[inline]
pub fn static_value<Input>() -> impl Parser<Input, Output = Value>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    static_value_()
}

#[test]
fn test_array() {
    assert_eq!(
        array().parse("[1,2,3,4,5]"),
        Ok((Value::literal_array(vec![1, 2, 3, 4, 5]), ""))
    );

    assert_eq!(
        array().parse("[1.4,2.231231,3.3123123,-12312.2,-2312.0]"),
        Ok((
            Value::literal_array(vec![1.4, 2.231231, 3.3123123, -12312.2, -2312.0]),
            ""
        ))
    );

    assert_eq!(
        array().parse("[]"),
        Ok((Value::literal_array::<Vec<i64>, i64>(vec![]), ""))
    );
}

#[test]
fn test_dict() {
    // TODO: create dictionary test
}
