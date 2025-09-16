package lexer

import (
	"testing"
)

func (k Keyword) String() string {
	switch k {
	case KBool:
		return "bool"
	case KInt:
		return "int"
	case KFloat:
		return "float"
	case KString:
		return "string"
	case KStruct:
		return "struct"
	case KTrue:
		return "true"
	case KFalse:
		return "false"
	case KIf:
		return "if"
	case KElse:
		return "else"
	case KWhile:
		return "while"
	case KBreak:
		return "break"
	case KContinue:
		return "continue"
	case KSwitch:
		return "switch"
	case KCase:
		return "case"
	case KDefault:
		return "default"
	default:
		return "unknown"
	}
}

func TestLexIdentifier(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `if test`,
			Tokens: []*Token{
				{Type: TKeyword, Value: KIf},
				{Type: TIdentifier, Value: "test"},
			},
		},
	})
}
