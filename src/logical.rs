use combine::{
    char::{char, spaces, string},
    error::ParseError,
    parser::Parser,
    stream::Stream,
};

//
// {{ if expr do }}
//
// #[inline]
// pub fn logical_stmt_if<I, P, A>(block: P, alt: A) -> impl Parser<Input = I, Output = ()>
// where
//     P: Parser<Input = I, Output = ()>,
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     string("{{").with(spaces().with(string("if")))
// }

//
// {{ elsif do }}
//
//
// #[inline]
// pub fn logical_stmt_alt<I, P, A>(block: P. alt: A) -> impl Parser<Input = I, Output = ()>
// where
//     P: Parser<Input = I, Output = ()>,
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
//     string("elsif").map(move |_| ())
// }

//
// {{ else do }}
//
// {{ end }}
//
// pub fn logical_stmt_else<I, P>() -> impl Parser<Input = I, Output = ()>
// where
//     P: Parser<Input = I, Output = ()>,
//     I: Stream<Item = char>,
//     I::Error: ParseError<I::Item, I::Range, I::Position>,
// {
// }
