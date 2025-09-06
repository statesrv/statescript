package lexer

import (
	"fmt"
)

// Error implements the error interface for errors encountered during lexing.
type Error struct {
	Context    string
	LineNumber int
	Message    string
}

func (e Error) Error() string {
	return fmt.Sprintf(
		"%s:%d - %s",
		e.Context,
		e.LineNumber,
		e.Message,
	)
}

func (l *Lexer) makeError(format string, a ...any) *Error {
	return &Error{
		Context:    l.context,
		LineNumber: l.lineNumber,
		Message:    fmt.Sprintf(format, a),
	}
}
