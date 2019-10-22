use std::{any::TypeId, error::Error as StdError, fmt};

#[derive(Debug, Clone, PartialEq)]
pub struct OperatorError<'a> {
    source: &'a str,
    target: TypeId,
}

//
// TODO: should I use T: fmt::Display instead?
//
impl<'a> fmt::Display for OperatorError<'a> {
    fn fmt(&self, fmt: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(
            fmt,
            "unexpected conversion error from str \"{:?}\" for Operator \"{:?}\"",
            self.source, self.target
        )
    }
}

impl<'a> StdError for OperatorError<'a> {
    #[inline]
    fn source(&self) -> Option<&(dyn StdError + 'static)> {
        None
    }
}

#[derive(Debug, Clone, PartialEq)]
pub enum ParseError<'a> {
    OperatorError(OperatorError<'a>),
}

impl<'a> ParseError<'a> {
    pub fn operator(source: &'a str, target: TypeId) -> Self {
        Self::OperatorError(OperatorError {
            source: source,
            target: target,
        })
    }
}

impl<'a> fmt::Display for ParseError<'a> {
    #[inline]
    fn fmt(&self, fmt: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            Self::OperatorError(e) => e.fmt(fmt),
        }
    }
}

// #[derive(Debug, Clone, PartialEq)]
// pub enum CastError<S>
// where
//     S: Sized + fmt::Debug,
// {
//     IncompatibleCast(S, TypeId),
// }

impl<'a> StdError for ParseError<'a> {
    #[inline]
    fn source(&self) -> Option<&(dyn StdError + 'static)> {
        None
    }
}
