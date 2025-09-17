package lexer

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	"github.com/nitroshare/compare"
)

type tTest struct {
	Name   string
	Input  string
	Tokens []*Token
	Error  bool
}

func doTests(t *testing.T, tests []*tTest) {
	for _, test := range tests {
		t.Run(test.Name, func(t *testing.T) {
			if test.Tokens == nil {
				test.Tokens = []*Token{}
			}
			r := strings.NewReader(test.Input)
			l, err := New(r, nil)
			if err != nil {
				compare.Compare(t, test.Error, true, true)
				return
			}
			tokens := []*Token{}
			for {
				n, err := l.Next()
				if err != nil {
					if !test.Error {
						t.Fatalf("unexpected error: %s", err)
					}
					break
				}
				if n.Type == TEOF {
					break
				}
				tokens = append(tokens, n)
			}
			if !reflect.DeepEqual(tokens, test.Tokens) {
				t.Fatalf("%v != %v", tokens, test.Tokens)
			}
		})
	}
}

type tBadReader struct{}

func (tBadReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test")
}

func TestFailedIO(t *testing.T) {
	r := &tBadReader{}
	_, err := New(r, nil)
	if err == nil {
		t.Fatal("error expected")
	}
}

func TestBadInput(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "InvalidChar",
			Input: "~",
			Error: true,
		},
	})
}

func TestSimpleProgram(t *testing.T) {
	doTests(t, []*tTest{
		{
			Name:  "Example",
			Input: `int i = 0`,
			Tokens: []*Token{
				{Type: TKeyword, Value: KInt},
				{Type: TIdentifier, Value: "i"},
				{Type: TDelimiter, Value: DAssignment},
				{Type: TInt, Value: int64(0)},
			},
		},
	})
}
