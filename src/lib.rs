#![feature(try_trait, const_fn, impl_trait_in_bindings)]
extern crate combine;
extern crate env_logger;
extern crate log;

pub mod ast;
pub mod bool;
pub mod error;
pub mod ident;
pub mod literal;
// pub mod logical;
pub mod expr;
pub mod operator;
pub mod types;
