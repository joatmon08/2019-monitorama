

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As an infrastructure engineer
// I want to know which resources belong to which environment
// So that I can correct add role-based access control.
func TestInfrastructureEnvironmentTags(t *testing.T) {
	hasEnvironmentTags := false
	assert.True(t, hasEnvironmentTags, "Found we do not have infrastructure environment differentiators")
}

// As an infrastructure engineer
// I want to know if a VM that I was supposed to patch still has vulnerabilities
// So that I can ensure it is secure.
func TestInfrastructureVMPatched(t *testing.T) {
	hasBeenPatched := false
	assert.True(t, hasBeenPatched, "Found VM with high vulnerabilities that has not been patched")
}

// As an infrastructure engineer
// I want to know if a change may be potentially disruptive
// So that I can ensure I maintain the upstream application SLOs.
func TestInfrastructureDryRun(t *testing.T) {
	hasDryRun := false
	assert.True(t, hasDryRun, "Found we need dry run capability describing disruption")
}
