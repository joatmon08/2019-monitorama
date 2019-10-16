package main

import (
	"testing"

	"github.com/joatmon08/slides/monitorama/mock"
	"github.com/stretchr/testify/assert"
)

// As a security engineer
// I want to ensure application dependencies are secure
// So that an attacker cannot exploit it.
func TestSecurityApplicationDependencyVulnerabilities(t *testing.T) {
	vulnerabilities := mock.ScanApplication()
	jsonBlob, err := mock.GenerateVulnerabilityReport(vulnerabilities)
	if err != nil {
		t.Error("Could not generate report")
	}
	t.Log(jsonBlob)
	assert.Equal(t, 0, len(vulnerabilities), "Found application dependencies with vulnerabilities")
}

// As a security engineer
// I want to ensure the application container is secure
// So that an attacker cannot exploit it.
func TestSecurityContainerVulnerabilities(t *testing.T) {
	vulnerabilities := mock.ScanContainer()
	jsonBlob, err := mock.GenerateVulnerabilityReport(vulnerabilities)
	if err != nil {
		t.Error("Could not generate report")
	}
	t.Log(jsonBlob)
	assert.Equal(t, 0, len(vulnerabilities), "Found container with vulnerabilities")
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

// As a security engineer
// I want to acknowledge 75% of critical alerts within one day
// So that I can assess if there is an attack.
func TestSecurityAlertsAcknowledged(t *testing.T) {
	acknowledged := 0
	alerts := mock.GetAlertsWithinLastDay(mock.TeamSecurity)
	for _, alert := range alerts {
		if alert.Acknowledged {
			acknowledged++
		}
	}
	acknowledgeRate := float64(acknowledged) / float64(len(alerts))
	assert.Greater(t, acknowledgeRate, 0.75, "Acknowledged fewer than 75% of alerts in past day")
}

// As a security engineer
// I want to know if a privilege escalation exploit attempt is made
// So that I can communicate and coordinate mitigation.
func TestSecurityAlertsForPrivilegeEscalationExploits(t *testing.T) {
	rules, err := mock.GetRulesByTag("privilege_escalation")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, len(rules), "Should have 1 rule for privilege escalation")
}

// As a security engineer
// I want to know there is an incident response plan
// So that I can properly identify, communicate, and mitigate a breach.
func TestSecurityIncidentResponsePlanAvailable(t *testing.T) {
	wiki := &mock.Wiki{}
	document, err := wiki.GetDocument("mock/incident_response.md")
	if err != nil {
		t.Error(err)
	}
	assert.Contains(t, document, "Incident Response Plan", "Did not find incident response plan as expected")
}
