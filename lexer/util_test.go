package lexer

import (
	"testing"
)

func TestEscapeChar(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `'\r' '\n' '\t' '\\'`,
			Tokens: []*Token{
				{Type: TChar, Value: '\r'},
				{Type: TChar, Value: '\n'},
				{Type: TChar, Value: '\t'},
				{Type: TChar, Value: '\\'},
			},
		},
		{
			Name:  "Unterminated",
			Input: `'\`,
			Error: true,
		},
	})
}

func TestWhitespace(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Skip",
			Input: " \t\r\n",
		},
	})
}
