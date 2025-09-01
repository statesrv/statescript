package lexer

// Keyword specifies a specific keyword.
type Keyword int

const (
	KBool Keyword = iota
	KInt
	KFloat
	KString
	KStruct

	KTrue
	KFalse

	KIf
	KElse

	KWhile
	KBreak
	KContinue

	KSwitch
	KCase
	KDefault
)
