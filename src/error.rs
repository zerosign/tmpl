// use combine::error::StreamError;
use std::{
    any::{Any, TypeId},
    fmt,
};

#[derive(Debug, Clone, PartialEq)]
pub enum ParseError {
    UnexpectedToken,
}

#[derive(Debug, Clone, PartialEq)]
pub enum CastError<S>
where
    S: Sized + fmt::Debug,
{
    IncompatibleCast(S, TypeId),
}
