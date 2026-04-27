# 013 正解例（課題別）

## 課題1: RESTハンドラ最小構成

```go
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func listProducts(w http.ResponseWriter, r *http.Request) {
	items := []Product{{ID: 1, Name: "book", Price: 1200}}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(items)
}
```

## 課題2: ページング・フィルタ

```go
func parseListQuery(r *http.Request) (limit, offset int, q string) {
	limit = 20
	offset = 0
	q = r.URL.Query().Get("q")
	if v := r.URL.Query().Get("limit"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n > 0 && n <= 100 {
			limit = n
		}
	}
	if v := r.URL.Query().Get("offset"); v != "" {
		if n, err := strconv.Atoi(v); err == nil && n >= 0 {
			offset = n
		}
	}
	return
}
```

## 課題3: バージョニング

```go
v1 := http.NewServeMux()
v1.HandleFunc("GET /products", listProducts)
root := http.NewServeMux()
root.Handle("/v1/", http.StripPrefix("/v1", v1))
```
