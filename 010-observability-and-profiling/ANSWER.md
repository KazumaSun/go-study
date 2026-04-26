# 010 正解例

```go
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	mux := http.NewServeMux()
	mux.HandleFunc("/work", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		time.Sleep(50 * time.Millisecond)
		logger.Info("request handled", "path", r.URL.Path, "duration_ms", time.Since(start).Milliseconds())
		w.WriteHeader(http.StatusOK)
	})

	_ = http.ListenAndServe(":8080", mux)
}
```
