### Known facts :

- lexer & parser for `tmpl` are using `rune` as basic type, means it supports utf8 natively

- `tmpl` didn't support adding an arbitary binary value into the template system (so don't
  include it into your template.

- escaping character are using `\` as defined in `token.SymbolEscape`.

- currently `tmpl` support minimum expression evaluator, please don't overexessively use it.

- it's inspired from `jinja` but it's (different from) not `jinja`.

# TOC

[Expression Notes](./expression.md)
[Function Notes](./function.md)
[Value Notes](./value.md)
[Editor Supports](./editor.md)
[Grammar](./grammar.bnf)
