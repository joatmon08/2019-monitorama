package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As an operator
// I want to know if the application is unhealthy
// So I can respond with a status notice.
func TestOperatorAlertApplicationUnhealthy(t *testing.T) {
	hasTriggered := false
	assert.True(t, hasTriggered, "Alert should have triggered for unhealthy application")
}

// As an operator
// I want to know who to escalate an issue
// So I can resolve a customer issue within my SLA.
func TestOperatorEscalationPath(t *testing.T) {
	hasEscalationPath := false
	assert.True(t, hasEscalationPath, "Found no escalation path defined")
}
