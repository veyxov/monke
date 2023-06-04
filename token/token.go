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
	semicolon = ";"

	lparen = "("
	rparen = ")"
	lbrace = "{"
	rbrace = "}"

	EOF      = "EOF"
	IDENT    = "IDENT"
	INT      = "INT"
	FUNCTION = "FUNCTION"
)
