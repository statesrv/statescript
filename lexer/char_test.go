package lexer

import (
	"testing"
)

func TestLexChar(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `'a'`,
			Tokens: []*Token{
				{Type: TChar, Value: 'a'},
			},
		},
		{
			Name:  "EOF",
			Input: `'`,
			Error: true,
		},
		{
			Name:  "Unterminated",
			Input: `'a`,
			Error: true,
		},
		{
			Name:  "Invalid",
			Input: `'ab'`,
			Error: true,
		},
		{
			Name:  "BadEscape",
			Input: `'\x'`,
			Error: true,
		},
	})
}
