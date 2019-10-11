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
	assert.True(t, isUpdated, "Found we have not updated our processing activity inventory in 30 days")
}

// As a compliance officer
// I want to ensure two people have reviewed a change to production
// So that we are SOX compliant.
func TestSOXReviewerCompliance(t *testing.T) {
	isUpdated := false
	assert.True(t, isUpdated, "Found we do not have two people reviewing a change")
}

// As a compliance officer
// I want to ensure we do not have PII data in logs
// So that we do not compromise customer data.
func TestPIIDataInLogs(t *testing.T) {
	hasPIIData := false
	assert.True(t, hasPIIData, "Found PII data in application logs")
}
