use combine::{
    parser::{
        char::{char, spaces},
        repeat,
    },
    unexpected_any, value, ParseError, Parser, Stream,
};

use std::iter;

#[derive(PartialEq, Default)]
pub(crate) struct StackSize {
    inner: usize,
}

impl StackSize {
    #[inline]
    pub const fn size(&self) -> usize {
        self.inner
    }
}

impl iter::Extend<char> for StackSize {
    fn extend<T>(&mut self, iter: T)
    where
        T: IntoIterator<Item = char>,
    {
        for _ in iter {
            self.inner += 1;
        }
    }
}

//
// Helper for creating (post)-padding space parser after actual parser.
//
//
#[inline]
pub fn lex<Input, P, O>(p: P) -> impl Parser<Input, Output = P::Output>
where
    P: Parser<Input, Output = O>,
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
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
pub fn closure<Input, P, O>(
    p: P,
    closure: &'static (char, char),
) -> impl Parser<Input, Output = P::Output>
where
    Input: Stream<Token = char>,
    P: Parser<Input, Output = O>,
    O: Clone,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    repeat::many::<StackSize, Input, _>(lex(char(closure.0)))
        .and(lex(p))
        .then(move |(stack, v)| {
            repeat::count::<StackSize, Input, _>(stack.size(), char(closure.1)).then(move |x| {
                if x == stack {
                    value(v.clone()).left()
                } else {
                    unexpected_any(closure.1).right()
                }
            })
        })
}

const ParaClosure: &'static (char, char) = &('(', ')');

#[inline]
pub fn para<Input, P, O>(p: P) -> impl Parser<Input, Output = P::Output>
where
    Input: Stream<Token = char>,
    P: Parser<Input, Output = O>,
    O: Clone,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    closure::<_, _, P::Output>(p, ParaClosure)
}

#[test]
fn test_lex() {
    assert_eq!(lex(char('a')).parse("a      "), Ok(('a', "")))
}

#[test]
fn test_closure() {
    use combine::char::string;

    assert_eq!(
        closure(string("true"), &('(', ')')).parse("(((true)))"),
        Ok(("true", ""))
    );

    assert!(closure(string("true"), &('(', ')'))
        .parse("(((true))")
        .is_err());

    assert_eq!(
        closure(string("true"), &('(', ')')).parse("true"),
        Ok(("true", ""))
    );
}
