package token

// String is useful for learning, but not performant. For better performance, int or byte can be used
type TokenType string

type Token struct {
    Type TokenType
    Literal string
}

const (
    ILLEGAL = "ILLEGAL"
    EOF = "EOF"

    // Identifiers + Literals
    IDENT = "IDENT" // Function_name, Variable_name etc..
    INT = "INT" // integer literals like 123135

    // Operators
    ASSIGN = "="
    PLUS = "+"
    MINUS = "-"
    MULTIPLY = "*"
    DIVIDE = "/"
    MOD = "%"
    BANG = "!"

    // Comparison
    LT = "<"
    GT = ">"
    EQ = "=="
    NOT_EQ = "!="
    LT_OR_EQ = "<="
    GT_OR_EQ = ">="

    // Delimiters
    COMMA = ","
    SEMICOLON = ";"
    LPAREN = "("
    RPAREN = ")"
    LBRACE = "{"
    RBRACE = "}"

    // Keywords
    FUNCTION = "FUNCTION"
    LET = "LET"
    RETURN = "RETURN"
    IF = "IF"
    ELSE = "ELSE"

    // Booleans
    TRUE = "TRUE"
    FALSE = "FALSE"
)

var keywords = map[string]TokenType {
    "function": FUNCTION,
    "let": LET,
    "return": RETURN,
    "if": IF,
    "else": ELSE,
    "true": TRUE,
    "false": FALSE,
}

func LookupIdent(ident string) TokenType {
    if tok, ok := keywords[ident]; ok {
        return tok
    }
    return IDENT
}
