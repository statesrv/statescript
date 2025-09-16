package lexer

// Type indicates the type of value a token contains.
type Type int

const (
	TEOF Type = iota
	TChar
	TString
	TInt
	TFloat
	TKeyword
	TIdentifier
	TDelimiter
)

// Token represents an individual lexeme extracted from the Lexer.
type Token struct {

	// Type indicates the type of token.
	Type Type

	// Value provides the relevant content of the token. Its type can be
	// asserted based on the value of Type and retreived with the convenience
	// methods.
	Value any
}

// AsRune returns the token's value when its type is TChar.
func (t Token) AsRune() rune {
	return t.Value.(rune)
}

// AsString returns the token's value when its type is TString or TIdentifier.
func (t Token) AsString() string {
	return t.Value.(string)
}

// AsInt64 returns the token's value when its type is TInt.
func (t Token) AsInt64() int64 {
	return t.Value.(int64)
}

// AsFloat64 returns the token's value when its type is TFloat.
func (t Token) AsFloat64() float64 {
	return t.Value.(float64)
}

// AsKeyword returns the token's value when its type is TKeyword.
func (t Token) AsKeyword() Keyword {
	return t.Value.(Keyword)
}

func (t Token) AsIdentifier() string {
	return t.Value.(string)
}

// AsDelimiter returns the token's value when its type is TDelimiter.
func (t Token) AsDelimiter() Delimiter {
	return t.Value.(Delimiter)
}
