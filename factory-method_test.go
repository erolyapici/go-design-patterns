package main

import (
	"testing"
)

func TestCreatePaymentMethodCash(t *testing.T) {

	payment, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("Cache must exist")
	}
	msg := payment.Pay(5)
	if msg != Cash {
		t.Errorf("Payment method wasn't correct")
	}
	t.Log(msg)
}
