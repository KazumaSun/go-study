# 003 正解例

```go
// service.go
package calc

import (
	"errors"
	"fmt"
)

var ErrDivideByZero = errors.New("divide by zero")

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide %d by %d: %w", a, b, ErrDivideByZero)
	}
	return a / b, nil
}
```

```go
// service_test.go
package calc

import (
	"errors"
	"testing"
)

func TestDivide(t *testing.T) {
	tests := []struct {
		name    string
		a, b    int
		want    int
		wantErr error
	}{
		{"ok", 8, 2, 4, nil},
		{"zero", 8, 0, 0, ErrDivideByZero},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Divide(tc.a, tc.b)
			if tc.wantErr != nil && !errors.Is(err, tc.wantErr) {
				t.Fatalf("expected %v, got %v", tc.wantErr, err)
			}
			if tc.wantErr == nil && got != tc.want {
				t.Fatalf("want %d, got %d", tc.want, got)
			}
		})
	}
}
```
