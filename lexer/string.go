package lexer

func (l *Lexer) lexString() *Token {
	v := ""
	for l.next(true) {
		if l.cur == '"' {
			break
		}
		c, ok := l.escapeChar()
		if !ok {
			return nil
		}
		v += string(c)
	}
	l.next(false)
	return &Token{
		Type:  TString,
		Value: v,
	}
}
