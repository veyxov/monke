package token

type TokenType string

type Token struct {
	TokenType TokenType
	Literal   string
}

const (
	ILLEGAL = "ILLEGAL"

	ASSIGN    = "="
	PLUS      = "+"
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	EOF      = "EOF"
	IDENT    = "IDENT"
	INT      = "INT"
	FUNCTION = "FUNCTION"
)
