package lexer

import (
	"fmt"
	"unicode"
)

type Lexer struct {
	input    string
	position int
	prevCh   rune // 前の文字
	ch       rune // 現在の文字
}

func NewLexer(input string) *Lexer {
	return &Lexer{
		input: input,
		position: 0,
		prevCh:   0,
		ch:       0,
	}
}

func (l *Lexer) Process() {
	for i, ch := range l.input {
		switch {
		case unicode.IsDigit(ch):
			fmt.Printf("position: %d, %c is digit\n", i, ch)
		case ch == '+':
			fmt.Printf("position: %d, %c is plus\n", i, ch)
		default:
			fmt.Printf("position: %d, %c is other\n", i, ch)
		}
	}
}

