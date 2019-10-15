//!
//! A list of  parser for ident like structure .
//!
//! it can be one of :
//! - macro call
//! - function call
//! - type declaration
//!
use crate::ast::Ident;
use combine::{
    many,
    parser::char::{digit, letter, lower, upper},
    token,
    ParseError, Parser, Stream,
};

//
// IdentLike = (Digit | Letter | "_")* "'"* .
//
#[inline]
pub fn ident_like<Input>() -> impl Parser<Input, Output = String>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    many(digit().or(letter()).or(token('_')))
        .and(many(token('\'')))
        .map(|(lhs, rhs): (String, String)| [lhs, rhs].concat())
}

//
// Ident = Lower IdentLike .
//
#[inline]
pub fn ident<Input>() -> impl Parser<Input, Output = Ident>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    lower::<Input>()
        .and(ident_like())
        .map(|(ch, mut b)| {
            b.insert(0, ch);
            b
        })
        .map(Ident::Ident)
}

#[inline]
pub fn macro_ident<Input>() -> impl Parser<Input, Output = Ident>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    lower::<Input>()
        .and(ident_like())
        .map(|(ch, mut b)| {
            b.insert(0, ch);
            b
        })
        .and(token('!'))
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
pub fn type_decl<Input>() -> impl Parser<Input, Output = Ident>
where
    Input: Stream<Token = char>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    upper::<Input>()
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
