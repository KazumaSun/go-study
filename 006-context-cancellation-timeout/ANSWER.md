# 006 正解例

```go
package main

import (
	"context"
	"fmt"
	"time"
)

func fetch(ctx context.Context) error {
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("done")
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := fetch(ctx); err != nil {
		fmt.Println("fetch error:", err)
	}
}
```
