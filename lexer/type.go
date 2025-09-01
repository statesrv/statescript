package lexer

// Type indicates the type of value a token contains.
type Type int

const (
	TEOF Type = iota
	TString
	TInt
	TFloat
	TKeyword
	TIdentifier
	TOperator
)
