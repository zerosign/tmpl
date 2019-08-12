# function

Function in tmpl can be extended by defining FunctionDecl directly or using `function.Convert` method.
Function in tmpl need to has several characteristics :

- function need to have at least 1 argument type variable outside `function.Context` and optional
  varargs arguments

- parameter with type `function.Context` need to be in the first parameter

- both `function.Context` and varargs arguments are optionals

- currently only support types that can be converted into `value.Kind` (didn't support custom struct in golang).

- return of the function need to be in size of 2, include the actual value (result, as lhs) of the function and
  error (as rhs).
