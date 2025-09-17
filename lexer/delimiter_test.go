package lexer

import (
	"testing"
)

func (d Delimiter) String() string {
	switch d {
	case DNone:
		return "none"
	case DNot:
		return "!"
	case DMultiply:
		return "*"
	case DMultiplyAssign:
		return "*="
	case DDivide:
		return "/"
	case DDivideAssign:
		return "/="
	case DModulo:
		return "%"
	case DModuloAssign:
		return "%="
	case DAdd:
		return "+"
	case DAddAssign:
		return "+="
	case DSubtract:
		return "-"
	case DSubtractAssign:
		return "-="
	case DLessThan:
		return "<"
	case DLessThanOrEqual:
		return "<="
	case DGreaterThan:
		return ">"
	case DGreaterThanOrEqual:
		return ">="
	case DEqual:
		return "="
	case DNotEqual:
		return "!="
	case DLogicalAnd:
		return "&&"
	case DLogicalOr:
		return "||"
	case DAssignment:
		return "="
	case DLBrace:
		return "{"
	case DRBrace:
		return "}"
	case DLParen:
		return "("
	case DRParen:
		return ")"
	case DLIndex:
		return "["
	case DRIndex:
		return "]"
	default:
		return "unknown"
	}
}

func TestLexDelimiter(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Good",
			Input: `+ +=`,
			Tokens: []*Token{
				{Type: TDelimiter, Value: DAdd},
				{Type: TDelimiter, Value: DAddAssign},
			},
		},
		{
			Name:  "Invalid",
			Input: `&`,
			Error: true,
		},
	})
}
