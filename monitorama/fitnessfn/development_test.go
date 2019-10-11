package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// As a developer
// I want to see my application's logs indexed
// So that I can find problems.
func TestDevelopmentApplicationLogging(t *testing.T) {
	inLogAggregator := false
	assert.True(t, inLogAggregator, "Found no logs in aggregator for application")
}

// As a developer
// I want to check my application's error rate is less than X% during updates
// So that I do not affect my users.
func TestDevelopmentApplicationErrorRate(t *testing.T) {
	errorRate := false
	assert.True(t, errorRate, "Found error rate is greater than X%")
}

// As a developer
// I want to check my application's code quality
// So that it is maintainable.
func TestDevelopmentApplicationMaintainabilityRating(t *testing.T) {
	maintainabilityRating := false
	assert.Greater(t, maintainabilityRating, 0.1, "Found maintainability rating less than 0.1 (B)")
}
