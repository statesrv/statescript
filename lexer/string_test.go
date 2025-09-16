package lexer

import (
	"testing"
)

func TestLexString(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `"" "test"`,
			Tokens: []*Token{
				{Type: TString, Value: ""},
				{Type: TString, Value: "test"},
			},
		},
		{
			Name:  "Unterminated",
			Input: `"`,
			Error: true,
		},
		{
			Name:  "BadEscape",
			Input: `"\x"`,
			Error: true,
		},
	})
}
