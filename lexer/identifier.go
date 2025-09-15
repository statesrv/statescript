package lexer

// Keyword specifies a specific keyword.
type Keyword int

const (
	KNone Keyword = iota
	KBool
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

func identifierToKeyword(v string) Keyword {
	switch v {
	case "bool":
		return KBool
	case "int":
		return KInt
	case "float":
		return KFloat
	case "string":
		return KString
	case "struct":
		return KStruct
	case "true":
		return KTrue
	case "false":
		return KFalse
	case "if":
		return KIf
	case "else":
		return KElse
	case "while":
		return KWhile
	case "break":
		return KBreak
	case "continue":
		return KContinue
	case "switch":
		return KSwitch
	case "case":
		return KCase
	case "default":
		return KDefault
	default:
		return KNone
	}
}

func (l *Lexer) lexIdentifier() *Token {
	v := string(l.cur)
	for l.next(false) {
		v += string(l.cur)
	}
	k := identifierToKeyword(v)
	if k != KNone {
		return &Token{
			Type:  TKeyword,
			Value: k,
		}
	}
	return &Token{
		Type:  TIdentifier,
		Value: v,
	}
}
