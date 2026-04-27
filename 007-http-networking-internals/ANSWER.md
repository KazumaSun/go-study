# 007 正解例（課題別）

## 課題1: カスタムhttp.Client

```go
package main

import (
	"net/http"
	"time"
)

func newClient() *http.Client {
	tr := &http.Transport{
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     30 * time.Second,
	}
	return &http.Client{
		Timeout:   3 * time.Second,
		Transport: tr,
	}
}
```

## 課題2: timeout設定済みサーバ

```go
package main

import (
	"net/http"
	"time"
)

func main() {
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      http.DefaultServeMux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	_ = srv.ListenAndServe()
}
```

## 課題3: 単純リトライ

```go
func retry3(do func() error) error {
	var err error
	wait := 100 * time.Millisecond
	for i := 0; i < 3; i++ {
		if err = do(); err == nil {
			return nil
		}
		time.Sleep(wait)
		wait *= 2
	}
	return err
}
```
