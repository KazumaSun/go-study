# 005 正解例

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	in := make(chan int)
	out := make(chan int)
	var wg sync.WaitGroup

	// fan-out
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for n := range in {
				out <- n * n
			}
		}()
	}

	go func() {
		for i := 1; i <= 5; i++ {
			in <- i
		}
		close(in)
		wg.Wait()
		close(out)
	}()

	// fan-in (single consumer)
	for v := range out {
		fmt.Println(v)
	}
}
```
