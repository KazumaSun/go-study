# 006 正解例（課題別）

## 課題1: timeout付き並列実行

```go
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func call(ctx context.Context, name string, d time.Duration) error {
	select {
	case <-time.After(d):
		fmt.Println(name, "ok")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	defer cancel()

	var wg sync.WaitGroup
	for _, job := range []struct {
		name string
		d    time.Duration
	}{
		{"a", 500 * time.Millisecond},
		{"b", 2 * time.Second},
	} {
		wg.Add(1)
		go func(name string, d time.Duration) {
			defer wg.Done()
			if err := call(ctx, name, d); err != nil {
				fmt.Println(name, "err:", err)
			}
		}(job.name, job.d)
	}
	wg.Wait()
}
```

## 課題2: 親キャンセルで子停止

```go
ctx, cancel := context.WithCancel(context.Background())
go worker(ctx)
time.Sleep(100 * time.Millisecond)
cancel()
```

## 課題3: request-idをcontext経由でログ

```go
type key string

const requestIDKey key = "rid"

ctx := context.WithValue(context.Background(), requestIDKey, "req-123")
rid, _ := ctx.Value(requestIDKey).(string)
fmt.Println("request_id=", rid)
```

## 口頭試問の回答例
- `context` を構造体フィールドに持たない（リクエスト寿命を超えやすい）
- timeoutは外側の呼び出し境界から決める
- valueは横断的メタデータの最小限に限定する
