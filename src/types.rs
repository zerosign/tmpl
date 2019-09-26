// in binary operator, both input should have the same
// types while the output could be different.
//
pub trait BinaryOp {
    type Input: Sized;
    type Output: Sized;

    fn apply(&self, lhs: Self::Input, rhs: Self::Input) -> Self::Output;
}

// Unary operator, this is required to implement NegationOp.
//
pub trait UnaryOp {
    type Input: Sized;
    type Output: Sized;

    fn apply(&self, input: Self::Input) -> Self::Output;
}
