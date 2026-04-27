# 002 正解例（課題別）

## 課題1: PaymentProcessor interface

```go
package payment

import "fmt"

type PaymentProcessor interface {
	Pay(amount int) error
}

type StripeProcessor struct{}

func (s *StripeProcessor) Pay(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %d", amount)
	}
	fmt.Printf("[stripe] paid=%d\n", amount)
	return nil
}
```

## 課題2: Repository interfaceのメモリ実装

```go
package todo

import "sync"

type Todo struct {
	ID   int
	Text string
}

type Repository interface {
	Save(t Todo) error
	FindAll() ([]Todo, error)
}

type MemoryRepo struct {
	mu   sync.RWMutex
	data []Todo
}

func (r *MemoryRepo) Save(t Todo) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.data = append(r.data, t)
	return nil
}

func (r *MemoryRepo) FindAll() ([]Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make([]Todo, len(r.data))
	copy(out, r.data)
	return out, nil
}
```

## 課題3: ログ差し替え可能設計

```go
package app

type Logger interface {
	Info(msg string)
}

type Service struct {
	log Logger
}

func NewService(log Logger) *Service {
	return &Service{log: log}
}

func (s *Service) Run() {
	s.log.Info("service started")
}
```

## 口頭試問の回答例
- 受け側が `interface` を定義すると依存方向を内側に向けられる
- pointer receiverは状態変更・コピー回避・一貫メソッドセットのために使う
- 巨大interfaceはモック肥大化と変更耐性の低下を招く
