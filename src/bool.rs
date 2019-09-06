use combine::{parser::char::string, ParseError, Parser, Stream};

//
// Parse boolean value into Value::Bool(_)
//
// Bool = "false" | "true" .
//
// ```rust
// assert_eq!(bool().parse("true").unwrap().0. Value::Bool(true));
// assert_eq!(bool().parse("false").unwrap().0. Value::Bool(false));
// ```
//
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

#[test]
fn bool_parse_test() {
    assert_eq!(bool().parse("true").unwrap().0, Value::Bool(true));
    assert_eq!(bool().parse("false").unwrap().0, Value::Bool(false));

    assert!(bool().parse("tru").is_err());
    assert!(bool().parse("fals").is_err());
}
