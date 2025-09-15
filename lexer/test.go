package lexer

func (l *Lexer) isAlphanum() bool {
	return l.cur == '_' ||
		l.cur >= 'A' && l.cur <= 'Z' ||
		l.cur >= 'a' && l.cur <= 'z' ||
		l.isDigit()
}

func (l *Lexer) isDigit() bool {
	return l.cur >= '0' && l.cur <= '9'
}
