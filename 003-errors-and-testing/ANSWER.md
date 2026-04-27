# 003 正解例（課題別）

## 課題1: エラー分類

```go
package user

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidEmail = errors.New("invalid email")

func ValidateEmail(email string) error {
	if !strings.Contains(email, "@") {
		return fmt.Errorf("email=%q: %w", email, ErrInvalidEmail)
	}
	return nil
}
```

## 課題2: table-driven test

```go
package user

import (
	"errors"
	"testing"
)

func TestValidateEmail(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr error
	}{
		{"ok", "a@example.com", nil},
		{"ng", "abc", ErrInvalidEmail},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ValidateEmail(tt.input)
			if tt.wantErr == nil && err != nil {
				t.Fatalf("unexpected err: %v", err)
			}
			if tt.wantErr != nil && !errors.Is(err, tt.wantErr) {
				t.Fatalf("want %v got %v", tt.wantErr, err)
			}
		})
	}
}
```

## 課題3: 境界値テスト

```go
package user

import "testing"

func TestValidateEmail_Boundary(t *testing.T) {
	cases := []string{"", "@", "a@", "@b.com"}
	for _, c := range cases {
		_ = ValidateEmail(c)
	}
}
```

## 口頭試問の回答例
- `%w` はエラー連鎖を保持、`%v` は文字列化のみ
- `errors.Is` はラップチェーン内に対象エラーがあれば true
- モック対象は「遅い/外部依存/不安定」な境界に限定する
