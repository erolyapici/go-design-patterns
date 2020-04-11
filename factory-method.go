package main

import (
	"errors"
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float32) int
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(method int) (PaymentMethod, error) {
	switch method {
	case Cash:
		return new(CachePaymentMethod), nil
	case DebitCard:
		return new(DebitCartPaymentMethod), nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not found", method))
	}
}

type CachePaymentMethod struct {
}
type DebitCartPaymentMethod struct {
}

func (c *CachePaymentMethod) Pay(amount float32) int {
	return Cash
}

func (c *DebitCartPaymentMethod) Pay(amount float32) int {
	return DebitCard
}
