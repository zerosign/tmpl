// in binary operator, both input should have the same
// types while the output could be different.
//
pub trait BinaryOp<I, O>
where
    I: Sized,
    O: Sized,
{
    fn apply(&self, lhs: I, rhs: I) -> O;
}

// Unary operator, this is required to implement NegationOp.
//
pub trait UnaryOp<I, O>
where
    I: Sized,
    O: Sized,
{
    fn apply(&self, input: I) -> O;
}

pub trait Apply<O>
where
    O: Sized,
{
    fn apply(&self) -> O;
}
