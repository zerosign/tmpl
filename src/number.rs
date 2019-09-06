use combine::{
    parser::{
        char::{digit, char},
        choice::optional,
        error::{unexpected_any},
        item::value,
        repeat::many,
    },
    ParseError, Parser, Stream
}

#[inline]
fn non_zero_digit<I>() -> impl Parser<Input = I, Output = char>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    digit().then(|d| {
        if d == '0' {
            unexpected_any(d)
                .message("shouldn't be `0` at beginning")
                .right()
        } else {
            value(d).left()
        }
    })
}

#[inline]
fn integer<I>() -> impl Parser<Input = I, Output = i64>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    char('0').map(|_| 0).or(non_zero_digit()
        .and(many(digit()))
        .then(|(d, etc): (char, String)| {
            let value = d as i64 - '0' as i64;
            etc.chars()
                .fold(value, |r, ch| (r * 10) + (ch as i64 - '0' as i64))
        }))
}

#[inline]
fn fractional<I>(exp: f64) -> impl Parser<Input = I, Output = f64>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    many(digit()).then(|mut digits: String| {
        digits.insert(0, '.');
        digits.insert(0, '0');
        match digits.parse::<f64>() {
            Ok(v) => value(exp + v).left(),
            Err(e) => unexpected_any(digits).right(),
        }
    })
}

#[inline]
fn number<I>() -> impl Parser<Input = I, Output = Number>
where
    I: Stream<Item = char>,
    I::Error: ParseError<I::Item, I::Range, I::Position>,
{
    optional(char('-'))
        .and(integer())
        .map(|(sign, n)| if sign.is_some() { -n } else { n })
        .and(optional(char('.')))
        .flat_map(|(exp, is_frac)| {
            if is_frac.is_some() {
                fractional(exp as f64).map(|v| Number::Double(v))
            } else {
                value(Number::Integer(exp))
            }
        })
}
