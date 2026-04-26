# 001 正解例

```go
package main

import (
	"fmt"
	"strconv"
)

const appName = "temp-converter"

func cToF(c float64) float64 { return c*9/5 + 32 }
func fToC(f float64) float64 { return (f - 32) * 5 / 9 }

func main() {
	fmt.Println(appName)

	c := 25.0
	f := cToF(c)
	fmt.Printf("%.1fC -> %.1fF\n", c, f)

	s := "86.0"
	parsed, err := strconv.ParseFloat(s, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%.1fF -> %.1fC\n", parsed, fToC(parsed))
}
```
