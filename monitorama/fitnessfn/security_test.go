package main

import (
	"testing"

	"github.com/joatmon08/slides/monitorama/mock"
	"github.com/stretchr/testify/assert"
)

// As a security engineer
// I want to ensure application dependencies are secure
// So that an attacker cannot exploit it.
func TestSecurityForApplicationDependencies(t *testing.T) {
	hasCleanScan := false
	assert.True(t, hasCleanScan, "Found application dependencies with vulnerabilities")
}

// As a security engineer
// I want to ensure the application container is secure
// So that an attacker cannot exploit it.
func TestSecurityForContainer(t *testing.T) {
	hasCleanScan := false
	assert.True(t, hasCleanScan, "Found application container with high vulnerabilities")
}

// As a security engineer
// I want to ensure the codebase does not have plaintext secrets
// So that our credentials are not compromised
func TestSecurityForPlaintextSecrets(t *testing.T) {
	hasNoSecrets := false
	assert.True(t, hasNoSecrets, "Found application codebase has plaintext secrets")
}

// As a security engineer
// I want to know how long someone uses root
// So that I can assess if it is a bad actor.
func TestSecurityForRootLogin(t *testing.T) {
	isRootLogin := false
	assert.True(t, isRootLogin, "Found no logs for root login")
}

// As a security engineer
// I want to know that a product owner and team lead have reviewed a change to production
// So that we are SOX compliant.
func TestSecuritySOXReviewWithProductOwnerAndTeamLead(t *testing.T) {
	pipeline := mock.Pipeline{}
	approvers := pipeline.GetApproversForPipeline()
	assert.NotEqual(t, approvers[0], approvers[1], "Approvers should not be the same")
	assert.Contains(t, []string{mock.RoleTeamLead, mock.RoleProductOwner}, approvers[0].Role, "Approver must be team lead or product owner")
	assert.Contains(t, []string{mock.RoleTeamLead, mock.RoleProductOwner}, approvers[1].Role, "Approver must be team lead or product owner")
}
