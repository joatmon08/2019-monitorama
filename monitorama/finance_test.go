package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As a finance stakeholder
// I want to ensure our bill does not exceed $100,000 a month
// So that we don't have sticker shock.
func TestFinanceCheckOurBill(t *testing.T) {
	inBudget := false
	assert.True(t, inBudget, "Found we have exceeded budget for the month")
}
