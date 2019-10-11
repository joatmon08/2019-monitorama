package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As a finance stakeholder
// I want to ensure our bill does not exceed $X a month
// So that we stay within budget.
func TestFinanceCheckOurBill(t *testing.T) {
	inBudget := false
	assert.True(t, inBudget, "Found we have exceeded budget for the month")
}

// As a finance stakeholder
// I want a way to examine my bill by line of business
// So that we can chargeback correctly.
func TestFinanceDetailedBill(t *testing.T) {
	hasFinanceTags := false
	assert.True(t, hasFinanceTags, "Found we cannot examine chargeback by line of business")
}
