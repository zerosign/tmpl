use combine::{
    char::{space, string},
    parser::range::take_until_range,
    ParseError, Parser, RangeStream,
};

use crate::ast::Comment;

//
// Parse block `comment` into
//
//
#[inline]
fn comment_block<'a, Input>() -> impl Parser<Input, Output = Comment<'a>>
where
    Input: RangeStream<Token = char, Range = &'a str>,
    Input::Error: ParseError<Input::Token, Input::Range, Input::Position>,
{
    // updates
    string("{#")
        .skip(space())
        .with(take_until_range("#}"))
        .skip(string("#}"))
        .map(|v: &str| Comment::new(v.trim_end_matches(' ')))
}

#[test]
fn test_comment_block() {
    // test empty comment
    assert!(comment_block().parse("{# #}").is_ok());

    // test general comment
    assert!(comment_block().parse("{# Hello world #}").is_ok());

    // test escaped comment
    assert!(comment_block().parse("{# {# #\\} #}").is_ok());
}
