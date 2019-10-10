package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As an operator
// I want to know if Consul has an issue
// So I can respond accordingly
func TestAlertForUnhealthyConsul(t *testing.T) {
	hasTriggered := false
	assert.True(t, hasTriggered, "Alert should have triggered for unhealthy Consul")
}