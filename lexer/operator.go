package lexer

// Operator indicates an operation to perform with one or more operands.
type Operator int

const (
	ONot Operator = iota
	OMultiply
	ODivide
	OModulo
	OAddOrPositive
	OSubtractOrNegative
	OLessThan
	OLessThanOrEqual
	OGreaterThan
	OGreaterThanOrEqual
	OEqual
	ONotEqual
	OAnd
	OOr
	OAssignment
	OLParen
	ORParen
)
