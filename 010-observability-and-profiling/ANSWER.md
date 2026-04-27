# 010 正解例（課題別）

## 課題1: 構造化ログ

```go
logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
logger.Info("request", "path", "/todos", "status", 200, "duration_ms", 12)
```

## 課題2: メトリクス出力

```go
var reqCount atomic.Int64

func handler(w http.ResponseWriter, r *http.Request) {
	reqCount.Add(1)
	w.WriteHeader(http.StatusOK)
}

func metrics(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "requests_total %d\n", reqCount.Load())
}
```

## 課題3: pprof導入

```go
import _ "net/http/pprof"

go func() {
	_ = http.ListenAndServe(":6060", nil)
}()
```

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=10
```
