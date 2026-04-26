# 004 正解例

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	worker := func(id int) {
		defer wg.Done()
		for j := range jobs {
			fmt.Printf("worker=%d job=%d\n", id, j)
		}
	}

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(i)
	}

	for j := 1; j <= 6; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()
}
```
