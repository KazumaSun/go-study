# 002 正解例

```go
package main

import "fmt"

type PaymentProcessor interface {
	Pay(amount int) error
}

type StripeProcessor struct{}

func (s *StripeProcessor) Pay(amount int) error {
	fmt.Printf("stripe paid: %d\n", amount)
	return nil
}

type CheckoutService struct {
	processor PaymentProcessor
}

func NewCheckoutService(p PaymentProcessor) *CheckoutService {
	return &CheckoutService{processor: p}
}

func (c *CheckoutService) Checkout(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("invalid amount: %d", amount)
	}
	return c.processor.Pay(amount)
}
```
