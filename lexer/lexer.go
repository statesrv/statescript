package lexer

import (
	"bufio"
	"io"
)

// Config provides configuration for the Lexer.
type Config struct {
	Context          string
	LineNumberOffset int
}

// Lexer takes input and produces a stream of tokens (lexemes).
type Lexer struct {
	r          *bufio.Reader
	context    string
	lineNumber int
	colNumber  int
	cur        rune
	err        *Error
}

// New creates a new Lexer instace for the provided io.Reader. The lexer is
// initialized with the values from the cfg parameter. If cfg is nil, default
// values are used.
func New(r io.Reader, cfg *Config) (*Lexer, *Error) {
	if cfg == nil {
		cfg = &Config{
			LineNumberOffset: 1,
		}
	}
	l := &Lexer{
		r:          bufio.NewReader(r),
		context:    cfg.Context,
		lineNumber: cfg.LineNumberOffset,
	}
	if l.context == "" {
		l.context = "[unknown]"
	}
	l.next(false)
	if l.err != nil {
		return nil, l.err
	}
	return l, nil
}

// Next returns the next token in sequence or an error if a syntax error is
// encoutered during the lexing process.
func (l *Lexer) Next() (*Token, *Error) {
	var t *Token
	if l.skipWhitespace() {
		switch {
		case l.cur == '\'':
			t = l.lexChar()
		case l.cur == '"':
			t = l.lexString()
		case l.isDigit():
			t = l.lexNumber()
		case l.isAlphanum():
			t = l.lexIdentifier()
		case l.isDelimiter():
			t = l.lexDelimiter()
		default:
			return nil, l.makeError("invalid character '%c'", l.cur)
		}
	}
	if t == nil && l.err == nil {
		t = &Token{
			Type: TEOF,
		}
	}
	return t, l.err
}
