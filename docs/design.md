Old design are being centered over tmpl grammar like any other template engine grammar
influenced by jinja or handlebars in general. However, I find it a bit limiting in term
of template evaluation. One example is when rendering yaml inside yaml template, one might
need context information of the tokens in last line to determine, how deep the placeholder
need to rendered to be considered as correct yaml. This means the function need to know
last rendered or unrendered parse tree node. This might be quite difficult to track if we
don't have full information of the parse tree that identically correlates to the template file.
This might includes omitted characters like whitespaces or unknown binary utfs.

Thus, I conclude my self that we need to have basic toplevel parser that

Parser

sample.yml.tmpl

Toplevel: yaml

```
test: quote(json_render(json_parse(data["content"])))
```

```
trait SyntaxTree {
    type Node;

    ...
}
```

```
enum YamlNode {
   YamlStr(String),
   YamlBool(bool),
   YamlInt(u64),
   YamlDouble(f64),
}
```

then

```
impl SyntaxTree for YamlParser {
    type Node = YamlNode;
    ...
}
```

quoting means lift the value to type of Toplevel string.
This means without quoting a node, there won't be a way
to render `JsonNode` inside `YamlNode`. With this we could design
safe rendering logic.

We also need to differentiate `SyntaxTree` & `ParseTree`. `SyntaxTree` should be used
when declaring & defining per format tree nodes while `ParseTree` are mostly related to
`TopLevel`/generic node. Thus, `SyntaxTree` node could reference to a segment in `ParseTree`
nodes.


Notes:
- I take insipiration on how `scalameta` models its tree internals as losseless nodes.
  see [ScalaMeta Syntax Tree](https://scalameta.org/docs/trees/guide.html).
  This give `scalameta` ability to deform & transform the nodes on each phase
  without losing any information to the original string.
- With loseless representations, we could even re-render all things including the comments
  inside the format without changing an inch of its definition content, since the tree
  are literaly a repr of the input file.
