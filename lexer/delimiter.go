package lexer

// Delimiter indicates an operator or separator.
type Delimiter int

const (
	DNone Delimiter = iota
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
	singleDelimiters = map[rune]Delimiter{
		'!': DNot,
		'*': DMultiply,
		'/': DDivide,
		'%': DModulo,
		'+': DAdd,
		'-': DSubtract,
		'<': DLessThan,
		'>': DGreaterThan,
		'=': DAssignment,
		'(': DLParen,
		')': DRParen,
		'[': DLIndex,
		']': DRIndex,
	}
	compoundDelimiters = map[string]Delimiter{
		"*=": DMultiplyAssign,
		"/=": DDivideAssign,
		"%=": DModuloAssign,
		"+=": DAddAssign,
		"-=": DSubtractAssign,
		"<=": DLessThanOrEqual,
		">=": DGreaterThanOrEqual,
		"==": DEqual,
		"!=": DNotEqual,
		"&&": DLogicalAnd,
		"||": DLogicalOr,
	}
)

func (l *Lexer) isDelimiter() bool {
	switch l.cur {
	case '!', '*', '/', '%',
		'+', '-', '<', '>',
		'=', '&', '|', '(',
		')', '[', ']':
		return true
	}
	return false
}

func (l *Lexer) lexDelimiter() *Token {
	v := l.cur
	if l.next(false) {
		d, ok := compoundDelimiters[string(v)+string(l.cur)]
		if ok {
			l.next(false)
			return &Token{
				Type:  TDelimiter,
				Value: d,
			}
		}
	}
	d, ok := singleDelimiters[v]
	if !ok {
		l.setError("invalid delimiter \"%c\"", v)
		return nil
	}
	return &Token{
		Type:  TDelimiter,
		Value: d,
	}
}
