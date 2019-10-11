package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As an consumer of the application
// I want to know if payment has successfully transacted within seconds
// So I can call more customers during the day.
func TestIfPaymentSuccessfulInSeconds(t *testing.T) {
	paymentSuccessful := false
	assert.True(t, paymentSuccessful, "Found payment did not completed in seconds")
}

// As an consumer of the application
// I want to know if the transaction has credited
// So I can buy my children's soccer team equipment this year.
func TestIfTransactionCreditedToSalesPerson(t *testing.T) {
	transactionCredited := false
	assert.True(t, transactionCredited, "Found transaction did not get credited to sales person")
}
