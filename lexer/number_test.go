package lexer

import (
	"testing"
)

func TestLexNumber(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `0 1. 1e1 1e-1 9223372036854775807 9223372036854775808`,
			Tokens: []*Token{
				{Type: TInt, Value: int64(0)},
				{Type: TFloat, Value: float64(1)},
				{Type: TFloat, Value: float64(10)},
				{Type: TFloat, Value: float64(0.1)},
				{Type: TInt, Value: int64(9223372036854775807)},
				{Type: TFloat, Value: float64(9223372036854775808)},
			},
		},
		{
			Name:  "BadFloat",
			Input: `1e`,
			Error: true,
		},
	})
}
