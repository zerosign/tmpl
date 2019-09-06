use combine::stream::state::SourcePosition;

#[derive(Debug, PartialEq)]
pub enum Ident {
    Ident(String),
    TypeDecl(String),
    MacroIdent(String),
}
