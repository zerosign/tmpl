use combine::{
    char::{space, string},
    parser::range::take_until_range,
    ParseError, Parser, Stream,
};

//
// Parse block `comment` into
//
//
#[inline]
fn comment_block<'a, I>() -> impl Parser<Input = I, Output = Comment<'a>>
where
    I: RangeStream<Item = char, Range = &'a str>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    // updates
    string("{#")
        .skip(space())
        .with(take_until_range("#}"))
        .skip(string("#}"))
        .map(|v: &str| Comment {
            value: v.trim_end_matches(' '),
        })
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
