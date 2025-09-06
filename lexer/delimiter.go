package lexer

// Delimiter indicates an operator or separator.
type Delimiter int

const (
	DCompound Delimiter = iota
	DNot

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

var (
	delimiterMap = map[rune]Delimiter{
		'!': DCompound,
		'*': DCompound,
		'/': DCompound,
		'%': DCompound,
		'+': DCompound,
		'-': DCompound,
		'<': DCompound,
		'>': DCompound,
		'=': DCompound,
		'&': DCompound,
		'|': DCompound,
		'(': DLParen,
		')': DRParen,
		'[': DLIndex,
		']': DRIndex,
	}
)
