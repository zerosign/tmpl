use combine::{
    parser::{
        char::{char, spaces},
        repeat,
    },
    unexpected_any, value, ParseError, Parser, Stream,
};

//
// Helper for creating (post)-padding space parser after actual parser.
//
//
#[inline]
pub fn lex<I, P, O>(p: P) -> impl Parser<Input = I, Output = P::Output>
where
    P: Parser<Input = I>,
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    p.skip(spaces())
}

//
// Closure helper for any closure like pattern that able to have recursive paren.
//
// This parser also checks whether left paren size are equal to right paren size.
//
// This parser are stacksafe.
//
#[inline]
pub fn closure<I, P, O>(
    p: P,
    closure: &'static (char, char),
) -> impl Parser<Input = I, Output = P::Output>
where
    I: Stream<Item = char>,
    O: Copy,
    P: Parser<Input = I, Output = O>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    repeat::many::<Vec<char>, _>(lex::<_, _, char>(char(closure.0)))
        .map(|v| v.len())
        .and(lex::<_, _, O>(p))
        .then(move |(stack, v)| {
            repeat::count::<Vec<char>, _>(stack, char(closure.1)).then(move |x| {
                if x.len() == stack {
                    value(v).left()
                } else {
                    unexpected_any(')').right()
                }
            })
        })
}

#[test]
fn test_lex() {
    assert_eq!(lex::<_, _, char>(char('a')).parse("a      "), Ok(('a', "")))
}

#[test]
fn test_closure() {
    use combine::char::string;

    assert_eq!(
        closure::<_, _, &str>(string("true"), &('(', ')')).parse("(((true)))"),
        Ok(("true", ""))
    );

    assert!(closure::<_, _, &str>(string("true"), &('(', ')'))
        .parse("(((true))")
        .is_err());
}
