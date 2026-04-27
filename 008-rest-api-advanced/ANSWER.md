# 008 正解例（課題別）

## 課題1: レイヤ分離の最小形

```go
// domain
type Todo struct {
	ID   int
	Text string
}

// repository
type TodoRepo interface {
	List() ([]Todo, error)
}

// usecase
type ListTodoUsecase struct{ repo TodoRepo }

func (u ListTodoUsecase) Execute() ([]Todo, error) {
	return u.repo.List()
}
```

## 課題2: 認証ミドルウェア（簡易）

```go
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
```

## 課題3: 一貫エラーレスポンス

```go
type APIError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

func writeError(w http.ResponseWriter, status int, code, msg string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(APIError{Code: code, Message: msg})
}
```
