use combine::{between, parser, parser::char::spaces, token, ParseError, Parser, Stream};

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

// #[inline]
// pub fn stackable_closure<Input, P, O>(
//     p: P,
//     closure: &'static (char, char),
// ) -> impl Parser<Input, Output = P::Output>
// where
//     Input: Stream<Token = char>,
//     P: Parser<Input, Output = O>,
//     O: Clone,
//     Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
// {

// }

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
    bracket: &'static (char, char),
) -> impl Parser<Input, Output = P::Output>
where
    Input: Stream<Token = char>,
    P: Parser<Input, Output = O>,
    O: Clone,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    closure_(p, bracket)
}

parser! {
    #[inline]
    pub fn closure_[Input, P, O](p: P, bracket: &'static (char, char))(Input) -> P::Output
    where [ Input: Stream<Token = char>, P: Parser<Input, Output = O>, O: Clone ] {
        between(lex(token(bracket.0)), lex(token(bracket.1)), closure(p, bracket))
    }
}

const ParaClosure: (char, char) = ('(', ')');

#[inline]
pub fn para<Input, P, O>(p: P) -> impl Parser<Input, Output = P::Output>
where
    Input: Stream<Token = char>,
    P: Parser<Input, Output = O>,
    O: Clone,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    closure::<_, _, P::Output>(p, &ParaClosure)
}

#[test]
fn test_lex() {
    assert_eq!(lex(token('a')).parse("a      "), Ok(('a', "")))
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
