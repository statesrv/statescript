package lexer

import (
	"fmt"
	"testing"
)

func (t Token) String() string {
	switch t.Type {
	case TEOF:
		return "EOF"
	case TChar:
		return fmt.Sprintf("[char '%c']", t.AsRune())
	case TString:
		return fmt.Sprintf("[string \"%s\"]", t.AsString())
	case TInt:
		return fmt.Sprintf("[int (%d)]", t.AsInt64())
	case TFloat:
		return fmt.Sprintf("[float (%.2f)]", t.AsFloat64())
	case TKeyword:
		return fmt.Sprintf("[keyword \"%s\"]", t.AsKeyword())
	case TIdentifier:
		return fmt.Sprintf("[identifier \"%s\"]", t.AsIdentifier())
	case TDelimiter:
		return fmt.Sprintf("[delimiter \"%s\"]", t.AsDelimiter())
	default:
		return "[unknown token]"
	}
}

func TestLookup(t *testing.T) {
	Token{Type: TChar, Value: 'a'}.AsRune()
	Token{Type: TString, Value: "a"}.AsString()
	Token{Type: TInt, Value: int64(1)}.AsInt64()
	Token{Type: TFloat, Value: float64(1.1)}.AsFloat64()
	Token{Type: TKeyword, Value: KBool}.AsKeyword()
	Token{Type: TIdentifier, Value: "a"}.AsIdentifier()
	Token{Type: TDelimiter, Value: DAdd}.AsDelimiter()
}
