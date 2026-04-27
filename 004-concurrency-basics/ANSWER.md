# 004 正解例（課題別）

## 課題1: 複数workerでジョブ処理

```go
package main

import (
	"fmt"
	"sync"
)

func main() {
	jobs := make(chan int)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("worker=%d job=%d\n", id, j)
			}
		}(w)
	}

	for i := 1; i <= 10; i++ {
		jobs <- i
	}
	close(jobs)
	wg.Wait()
}
```

## 課題2: timeout付き処理

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(2 * time.Second)
		close(done)
	}()

	select {
	case <-done:
		fmt.Println("finished")
	case <-time.After(1 * time.Second):
		fmt.Println("timeout")
	}
}
```

## 課題3: デッドロック例（悪い例）

```go
package main

func main() {
	ch := make(chan int)
	ch <- 1 // 受信側がいないのでここで停止
}
```

## 口頭試問の回答例
- closeは送信側が行う
- buffered channelは一時的な吸収には有効だが根本的な遅さは隠す
- goroutineリークは終了条件未設計で発生する
