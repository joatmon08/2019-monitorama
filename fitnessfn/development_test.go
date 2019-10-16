package main

import (
	"testing"
	"time"

	"github.com/joatmon08/slides/monitorama/mock"
	"github.com/stretchr/testify/assert"
)

// As a developer
// I want to push to production in 10 minutes
// So that I can release features more quickly.
func TestDevelopmentPipelineDeploymentTime(t *testing.T) {
	pipeline := mock.Pipeline{
		StartTime: time.Now(),
		EndTime:   time.Now().Add(time.Duration(12) * time.Minute),
	}
	deploymentTime := pipeline.GetDeploymentTime()
	assert.Less(t, deploymentTime.Minutes(), float64(10), "Deployment time should be less than 10 minutes")
}
