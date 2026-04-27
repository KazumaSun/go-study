package main

import (
	"fmt"

	"002-struct-interface-design/cmd/payment"
	"002-struct-interface-design/cmd/todo"
)

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

type StdLogger struct{}

func (l *StdLogger) Info(msg string) {
	fmt.Println(msg)
}

type PrefixLogger struct {
	prefix string
}

func (l *PrefixLogger) Info(msg string) {
	fmt.Printf("%s %s\n", l.prefix, msg)
}

func main() {
	svc := NewService(&StdLogger{})
	svc.Run()

	processor := &payment.StripeProcessor{}
	if err := processor.Pay(1000); err != nil {
		svc.log.Info(fmt.Sprintf("payment error: %v", err))
	}

	repo := &todo.MemoryRepo{}
	_ = repo.Save(todo.Todo{ID: 1, Text: "study interface"})
	_ = repo.Save(todo.Todo{ID: 2, Text: "study repository"})

	items, err := repo.FindAll()
	if err != nil {
		svc.log.Info(fmt.Sprintf("todo error: %v", err))
		return
	}
	svc.log.Info(fmt.Sprintf("saved todos: %d", len(items)))

	prefixSvc := NewService(&PrefixLogger{prefix: "[APP]"})
	prefixSvc.Run()
}
