# konnnichiwa


## まずはgoでhello worldする

参考: http://go.dev/doc/tutorial/getting-started

```bash
# goのversionを確認する
go version
# => go version go1.25.0 darwin/arm64

# モジュールの初期化
go mod init new-lang
```
hello.go

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
hello.goを実行する
```bash
go run .
#=> Hello, World!

```

## プログラミング言語panの誕生
名前：pan
拡張子：pan
アイコン：🍞
