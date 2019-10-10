package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As a GDPR Data Controller
// I want to ensure our inventory of processing activities has been updated
// So that we are in compliance.
func TestGDPRProcessingActivityInventoryUpdated(t *testing.T) {
	isUpdated := false
	assert.True(t, isUpdated, "Found we have not updated our processing activity inventory")
}
