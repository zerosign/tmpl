use combine::error::StreamError;
use std::convert::Into;

#[derive(Debug, Clone, PartialEq)]
pub enum ParseError {
    UnexpectedToken,
}
