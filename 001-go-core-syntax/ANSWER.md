# 001 正解例（課題別）

## 課題1: 温度変換CLI

```go
package main

import (
	"flag"
	"fmt"
)

func cToF(c float64) float64 { return c*9/5 + 32 }
func fToC(f float64) float64 { return (f - 32) * 5 / 9 }

func main() {
	mode := flag.String("mode", "c2f", "c2f or f2c")
	value := flag.Float64("value", 0, "temperature value")
	flag.Parse()

	switch *mode {
	case "c2f":
		fmt.Printf("%.2fC => %.2fF\n", *value, cToF(*value))
	case "f2c":
		fmt.Printf("%.2fF => %.2fC\n", *value, fToC(*value))
	default:
		fmt.Println("mode must be c2f or f2c")
	}
}
```

## 課題2: 文字列解析CLI

```go
package main

import (
	"fmt"
	"os"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: app <text>")
		return
	}
	s := os.Args[1]

	letters, digits, spaces := 0, 0, 0
	for _, r := range s {
		switch {
		case unicode.IsLetter(r):
			letters++
		case unicode.IsDigit(r):
			digits++
		case unicode.IsSpace(r):
			spaces++
		}
	}

	fmt.Printf("len=%d letters=%d digits=%d spaces=%d\n", len([]rune(s)), letters, digits, spaces)
}
```

## 課題3: defer実行順の確認

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("defer-1")
	defer fmt.Println("defer-2")
	defer fmt.Println("defer-3")
	fmt.Println("end")
}
```

## 口頭試問の回答例
- `:=` が使えない場所: パッケージスコープ
- `defer` の評価タイミング: 引数評価は `defer` 記述時、実行は関数return直前
- `while` が不要な理由: `for` で条件式だけ書けば同等表現が可能
