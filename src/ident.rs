use crate::ast::Ident;
use combine::{
    error::ParseError,
    parser::{
        char::{char, digit, letter, lower, upper},
        repeat::many,
    },
    Parser, Stream,
};

//
// IdentLike = (Digit | Letter | "_")* "'"* .
//
#[inline]
pub fn ident_like<I>() -> impl Parser<Input = I, Output = String>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    many(digit().or(letter()).or(char('_')))
        .and(many(char('\'')))
        .map(|(lhs, rhs): (String, String)| [lhs, rhs].concat())
}

//
// Ident = Lower IdentLike .
//
#[inline]
pub fn ident<I>() -> impl Parser<Input = I, Output = Ident>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    lower()
        .and(ident_like())
        .map(|(ch, mut b)| {
            b.insert(0, ch);
            b
        })
        .map(Ident::Ident)
}

#[inline]
pub fn macro_ident<I>() -> impl Parser<Input = I, Output = Ident>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    lower()
        .and(ident_like())
        .map(|(ch, mut b)| {
            b.insert(0, ch);
            b
        })
        .and(char('!'))
        .map(|(mut b, ch)| {
            b.push(ch);
            b
        })
        .map(Ident::MacroIdent)
}

//
// TypeDecl = Upper IdentLike .
//
#[inline]
pub fn type_decl<I>() -> impl Parser<Input = I, Output = Ident>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    upper()
        .and(ident_like())
        .map(|(ch, mut b)| {
            b.insert(0, ch);
            b
        })
        .map(Ident::TypeDecl)
}

#[test]
fn test_type_decl() {
    // allow char '_'
    assert_eq!(
        type_decl().parse("Hello_World").unwrap().0,
        Ident::TypeDecl(String::from("Hello_World"))
    );

    // allow char '\''
    assert_eq!(
        type_decl().parse("Integer'").unwrap().0,
        Ident::TypeDecl(String::from("Integer'"))
    );
}

#[test]
fn test_ident() {
    assert_eq!(
        ident().parse("hello_world").unwrap().0,
        Ident::Ident(String::from("hello_world")),
    );
}

#[test]
fn test_macro_ident() {
    assert_eq!(
        macro_ident().parse("println!").unwrap().0,
        Ident::MacroIdent(String::from("println!")),
    )
}
