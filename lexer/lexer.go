package lexer

import (
	"fmt"
	"unicode"
)

func Lexer(input string) {
	fmt.Printf("lexer input: %s\n", input)

	for i, ch := range input {
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

