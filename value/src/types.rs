use std::collections::HashMap;

//
// Number representations in `tmpl`.
//
#[derive(Debug, Clone, PartialEq)]
pub enum Number {
    Integer(i64),
    Double(f64),
}

// Most literal types.
//
// - Number ~ f64 and i64 (no unsigned integer)
// - String ~ String (owned type)
// - Bool ~ boolean type
// - Optional type
//
#[derive(Debug, Clone, PartialEq)]
pub enum Literal {
    Number(Number),
    String(String),
    Bool(bool),
    None,
}

impl Literal {
    #[inline]
    pub(crate) fn integer<V>(v: V) -> Literal
    where
        V: Into<i64>,
    {
        Literal::Number(Number::Integer(v.into()))
    }

    #[inline]
    pub(crate) fn double<V>(v: V) -> Literal
    where
        V: Into<f64>,
    {
        Literal::Number(Number::Double(v.into()))
    }

    #[inline]
    pub(crate) fn string<V>(v: V) -> Literal
    where
        V: Into<String>,
    {
        Literal::String(v.into())
    }

    #[inline]
    pub(crate) fn bool<V>(v: V) -> Literal
    where
        V: Into<bool>,
    {
        Literal::Bool(v.into())
    }
}

macro_rules! literal_conv {
    ($($conv:path => [$($src:ty),*]),*) => {
        $($(impl From<$src> for Literal {

            #[inline]
            fn from(v: $src) -> Self {
                $conv(v)
            }
        })*)*
    }
}

//
// Note: we don't support u64 at this point.
//
literal_conv!(
    Literal::integer => [u8, u16, u32, i8, i16, i32, i64],
    Literal::double  => [f32, f64],
    Literal::string  => [String, &'static str],
    Literal::bool    => [bool]
);

// value =
//   value = literal
//   dictionary<string, value>
//   array<value>
//
// - Dictionary
// - Array
//
#[derive(Debug, PartialEq)]
pub enum Value {
    Literal(Literal),
    Dictionary(HashMap<String, Value>),
    Array(Vec<Value>),
}

impl Value {
    pub fn literal_array<I, V>(i: I) -> Value
    where
        I: IntoIterator<Item = V>,
        V: Into<Literal>,
    {
        Value::Array(
            i.into_iter()
                .map(|v| Value::Literal(V::into(v)))
                .collect::<Vec<Value>>(),
        )
    }
}

// @TODO: @zerosign array!
//
// convert any primitive values in array declaration into
// literal values automatically
// ```
// assert_eq!(array!([1, 2, [2]]), Value::Array([
//    Literal::Number(Number::Integer(1)),
//    Literal::Number(Number::Integer(1)),
//    Value::Array([
//       Literal::Number(Number::Integer(2)),
//    ]),
// ]));
// ```
//
macro_rules! array {

}

// TODO: @zerosign dict!
//
// convert any primitive values including array to value literal directly.
// ```
// assert_eq!(dict!(
//
// ))
// ```
macro_rules! dict {

}
