# 005 正解例（課題別）

## 課題1: worker pool

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int, 20)
	results := make(chan int, 20)
	var wg sync.WaitGroup

	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				results <- j * j
			}
		}()
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
	close(results)

	for r := range results {
		fmt.Println(r)
	}
}
```

## 課題2: fan-out / fan-in

```go
package main

import "fmt"

func main() {
	src := make(chan int)
	a := make(chan int)
	b := make(chan int)
	merged := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			src <- i
		}
		close(src)
	}()
	go func() { for v := range src { a <- v * 2 }; close(a) }()
	go func() { for v := range src { b <- v * 3 }; close(b) }()
	go func() { for v := range a { merged <- v }; for v := range b { merged <- v }; close(merged) }()

	for v := range merged {
		fmt.Println(v)
	}
}
```

## 課題3: race再現と修正

```go
// bad
counter := 0
go func() { counter++ }()
go func() { counter++ }()

// good
var mu sync.Mutex
go func() { mu.Lock(); counter++; mu.Unlock() }()
go func() { mu.Lock(); counter++; mu.Unlock() }()
```
