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

var (
	keywords = map[string]Keyword{
		"bool":     KBool,
		"int":      KInt,
		"float":    KFloat,
		"string":   KString,
		"struct":   KStruct,
		"true":     KTrue,
		"false":    KFalse,
		"if":       KIf,
		"else":     KElse,
		"while":    KWhile,
		"break":    KBreak,
		"continue": KContinue,
		"switch":   KSwitch,
		"case":     KCase,
		"default":  KDefault,
	}
)

func (l *Lexer) lexIdentifier() *Token {
	v := string(l.cur)
	for l.next(false) {
		if !l.isAlphanum() {
			break
		}
		v += string(l.cur)
	}
	k, ok := keywords[v]
	if ok {
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
