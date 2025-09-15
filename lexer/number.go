package lexer

import (
	"strconv"
)

func (l *Lexer) lexNumberSegment() (string, bool) {
	v := ""
	for l.cur >= '0' && l.cur <= '9' {
		v += string(l.cur)
		if !l.next(false) {
			return v, false
		}
	}
	return v, true
}

func (l *Lexer) lexNumberSegmentOptional() string {
	v, _ := l.lexNumberSegment()
	return v
}

func (l *Lexer) lexNumber() *Token {
	v, ok := l.lexNumberSegment()
	if !ok || l.cur != '.' && l.cur != 'e' {
		i, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return &Token{
				Type:  TInt,
				Value: i,
			}
		}
	}
	if l.cur == '.' {
		v += string(l.cur)
		l.next(false)
	}
	v += l.lexNumberSegmentOptional()
	if l.cur == 'e' || l.cur == 'E' {
		v += string(l.cur)
	}
	if l.cur == '+' || l.cur == '-' {
		v += string(l.cur)
	}
	v += l.lexNumberSegmentOptional()
	f, err := strconv.ParseFloat(v, 64)
	if err != nil {
		l.setError("incorrectly formatted float")
		return nil
	}
	return &Token{
		Type:  TFloat,
		Value: f,
	}
}
