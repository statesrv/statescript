package lexer

func (l *Lexer) lexChar() *Token {
	if !l.next(true) {
		return nil
	}
	v, ok := l.escapeChar()
	if !ok {
		return nil
	}
	if !l.next(true) {
		return nil
	}
	if l.cur != '\'' {
		l.setError("expected \"'\"")
		return nil
	}
	return &Token{
		Type:  TChar,
		Value: v,
	}
}
