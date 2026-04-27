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
