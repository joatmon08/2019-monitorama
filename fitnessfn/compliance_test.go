package main

import (
	"testing"

	"github.com/joatmon08/slides/monitorama/mock"
	"github.com/stretchr/testify/assert"
)

// As a compliance officer
// I want to ensure two people have reviewed a change to production
// So that we are SOX compliant.
func TestComplianceSOXReviewerCompliance(t *testing.T) {
	p := mock.Pipeline{}
	approvers := p.GetApproversForPipeline()
	assert.Equal(t, len(approvers), 2, "Should have two approvers")
	assert.NotEqual(t, approvers[0], approvers[1], "Approvers should not be the same")
}