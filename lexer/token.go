package lexer

// Token represents an individual lexeme extracted from the Lexer.
type Token struct {

	// Type indicates the type of token.
	Type Type

	// Value provides the relevant content of the token. Its type can be
	// asserted based on the value of Type and retreived with the convenience
	// methods.
	Value any
}

// String returns the token's value when its type is TString or TIdentifier.
func (t Token) String() string {
	return t.Value.(string)
}

// Int64 returns the token's value when its type is TInt.
func (t Token) Int64() int64 {
	return t.Value.(int64)
}

// Float64 returns the token's value when its type is TFloat.
func (t Token) Float64() float64 {
	return t.Value.(float64)
}

// Keyword returns the token's value when its type is TKeyword.
func (t Token) Keyword() Keyword {
	return t.Value.(Keyword)
}

// Delimiter returns the token's value when its type is TDelimiter.
func (t Token) Delimiter() Delimiter {
	return t.Value.(Delimiter)
}
