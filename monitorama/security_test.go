package main

import (
	"testing"

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
// I want to ensure the VM hosting an application is secure
// So that an attacker cannot exploit it.
func TestSecurityForVM(t *testing.T) {
	hasCleanScan := false
	assert.True(t, hasCleanScan, "Found VM with high vulnerabilities")
}

// As a security engineer
// I want to ensure the codebase does not have plaintext secrets
// So that our credentials are not compromised
func TestSecurityForPlaintextSecrets(t *testing.T) {
	hasNoSecrets := false
	assert.True(t, hasNoSecrets, "Found application codebase has plaintext secrets")
}

// As a security engineer
// I want to know when someone logs in as root
// So that I can assess if it is a bad actor.
func TestSecurityForRootLogin(t *testing.T) {
	isRootLogin := false
	assert.True(t, isRootLogin, "Found no logs for root login")
}

// As a security engineer
// I want to know that a product owner and team lead have reviewed a change to production
// So that we are SOX compliant.
func TestSOXReviewWithProductOwnerAndTeamLead(t *testing.T) {
	hasRolesReviewed := false
	assert.True(t, hasRolesReviewed, "Found we do not have a product owner and/or team lead reviewing")
}
