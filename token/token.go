package token

type TokenType string

type Token struct {
	TokenType TokenType
	Literal   string
}
