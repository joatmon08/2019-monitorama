package main

import (
	"testing"

	"github.com/joatmon08/slides/monitorama/mock"
	"github.com/stretchr/testify/assert"
)

// As a GDPR Data Controller
// I want to ensure our inventory of processing activities has been updated
// So that we are in compliance.
func TestComplianceGDPRProcessingActivityInventoryUpdated(t *testing.T) {
	isUpdated := false
	assert.True(t, isUpdated, "Found we have not updated our processing activity inventory in 30 days")
}

// As a compliance officer
// I want to ensure two people have reviewed a change to production
// So that we are SOX compliant.
func TestComplianceSOXReviewerCompliance(t *testing.T) {
	p := mock.Pipeline{}
	approvers := p.GetApproversForPipeline()
	assert.Equal(t, len(approvers), 2, "Should have two approvers")
	assert.NotEqual(t, approvers[0], approvers[1], "Approvers should not be the same")
}

// As a compliance officer
// I want to ensure we do not have PII data in logs
// So that we do not compromise customer data.
func TestCompliancePIIDataInLogs(t *testing.T) {
	hasPIIData := false
	assert.True(t, hasPIIData, "Found PII data in application logs")
}
