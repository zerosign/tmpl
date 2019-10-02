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
    O: Clone,
    P: Parser<Input = I, Output = O>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    repeat::many::<StackSize, _>(lex::<_, _, char>(char(closure.0)))
        .and(lex::<_, _, O>(p))
        .then(move |(stack, v)| {
            repeat::count::<StackSize, _>(stack.size(), char(closure.1)).then(move |x| {
                if x == stack {
                    value(v.clone()).left()
                } else {
                    unexpected_any(')').right()
                }
            })
        })
}

const ParaClosure: &'static (char, char) = &('(', ')');

#[inline]
pub fn para<I, P, O>(p: P) -> impl Parser<Input = I, Output = P::Output>
where
    I: Stream<Item = char>,
    O: Clone,
    P: Parser<Input = I, Output = O>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    closure(p, ParaClosure)
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

    assert_eq!(
        closure::<_, _, &str>(string("true"), &('(', ')')).parse("true"),
        Ok(("true", ""))
    );
}
