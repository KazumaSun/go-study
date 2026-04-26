# 008 正解例

```go
package main

import (
	"encoding/json"
	"net/http"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}
		_ = json.NewEncoder(w).Encode([]Todo{{ID: 1, Text: "study go"}})
	})
	_ = http.ListenAndServe(":8080", mux)
}
```
