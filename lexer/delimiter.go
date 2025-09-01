package lexer

// Delimiter indicates an operator or separator.
type Delimiter int

const (
	DNot Delimiter = iota

	DMultiply
	DMultiplyAssign
	DDivide
	DDivideAssign
	DModulo
	DModuloAssign
	DAdd
	DAddAssign
	DSubtract
	DSubtractAssign

	DLessThan
	DLessThanOrEqual
	DGreaterThan
	DGreaterThanOrEqual
	DEqual
	DNotEqual

	DLogicalAnd
	DLogicalOr

	DAssignment

	DLParen
	DRParen
	DLIndex
	DRIndex
)
