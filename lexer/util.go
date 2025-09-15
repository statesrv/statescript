package lexer

import (
	"io"
)

func (l *Lexer) escapeChar() (rune, bool) {
	v := l.cur
	if v == '\\' {
		if !l.next(true) {
			return 0, false
		}
		switch l.cur {
		case 'r':
			v = '\r'
		case 'n':
			v = '\n'
		case 't':
			v = '\t'
		case '\\':
			v = '\\'
		default:
			l.setError("invalid escape sequence '\\%c'", l.cur)
			return 0, false
		}
	}
	return v, true
}

func (l *Lexer) next(must bool) bool {
	r, _, err := l.r.ReadRune()
	l.cur = r
	if err != nil {
		if err != io.EOF {
			l.setError("%s", err.Error())
		} else if must {
			l.setError("unexpected end-of-file")
		}
		return false
	}
	l.colNumber++
	return true
}

func (l *Lexer) skipWhitespace() bool {
	for l.cur == ' ' ||
		l.cur == '\t' ||
		l.cur == '\r' ||
		l.cur == '\n' {
		if l.cur == '\n' {
			l.lineNumber++
			l.colNumber = 0
		}
		if !l.next(false) {
			return false
		}
	}
	return true
}
