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
    pub fn integer<V>(v: V) -> Literal
    where
        V: Into<i64>,
    {
        Literal::Number(Number::Integer(v.into()))
    }

    #[inline]
    pub fn double<V>(v: V) -> Literal
    where
        V: Into<f64>,
    {
        Literal::Number(Number::Double(v.into()))
    }

    #[inline]
    pub fn string<V>(v: V) -> Literal
    where
        V: Into<String>,
    {
        Literal::String(v.into())
    }

    #[inline]
    pub fn bool<V>(v: V) -> Literal
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

macro_rules! array {
    [] => (Value::Array(Vec::<Value>::new()));
    [$($val:expr),*] => (Value::Array(<[_]>::into_vec(Box::new([$(Value::from($val)),*]))));
}

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
    #[inline]
    pub fn integer<V>(v: V) -> Value
    where
        V: Into<i64>,
    {
        Value::Literal(Literal::integer(v))
    }

    #[inline]
    pub fn double<V>(v: V) -> Value
    where
        V: Into<f64>,
    {
        Value::Literal(Literal::double(v))
    }

    #[inline]
    pub fn string<V>(v: V) -> Value
    where
        V: Into<String>,
    {
        Value::Literal(Literal::string(v))
    }

    #[inline]
    pub fn dict() -> Value {
        Value::Dictionary(HashMap::new())
    }

    #[inline]
    pub fn list() -> Value {
        Value::Array(vec![])
    }

    #[inline]
    pub fn bool<V>(v: V) -> Value
    where
        V: Into<bool>,
    {
        Value::Literal(Literal::bool(v))
    }
}

macro_rules! value_conv {
    ($($conv:path => [$($src:ty),*]),*) => {
        $($(impl From<$src> for Value {

            #[inline]
            fn from(v: $src) -> Self {
                $conv(v)
            }
        })*)*
    }
}

value_conv!(
    Value::integer => [u8, u16, u32, i8, i16, i32, i64],
    Value::double  => [f32, f64],
    Value::string  => [String, &'static str],
    Value::bool    => [bool]
);

#[test]
fn test_macro_rule_empty_array() {
    assert_eq!(array![], Value::Array(vec![]));
}

#[test]
fn test_macro_rule_literal_array() {
    assert_eq!(
        array![1, 2, 3.2, 4, "test"],
        Value::Array(vec![
            Value::integer(1),
            Value::integer(2),
            Value::double(3.2),
            Value::integer(4),
            Value::string("test"),
        ])
    );
}

#[test]
fn test_macro_rule_complex_array() {
    assert_eq!(
        array![1, array![1, 2]],
        Value::Array(vec![
            Value::integer(1),
            Value::Array(vec![Value::integer(1), Value::integer(2)])
        ])
    );
}
