package main

import (
	"fmt"
	"new-lang/lexer"
	"os"
	"log"
	"io"
)

func main() {
	// debug print
	fmt.Printf("hello")

	// panファイルを読み込みたい
	file, err := os.Open("pan.pan")
	if err != nil {
		log.Fatal(err)
	}
	
	// ファイルの中を読み取る
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("pan.pan content: %s\n", string(content))

	// ファイルとじる
	defer file.Close()

	// lexerを呼び出す
	lexer.Lexer()
}
