package checker

import (
)

type expectation struct {
	expectedValue string
	expectedCount int
	actualCount int
}

func newExpectation(val string, cnt int) *expectation {
	return &expectation{
		expectedValue: val,
		expectedCount: cnt,
	}
}

func (e *expectation) matches(val string) bool {
	if e.actualCount >= e.expectedCount {
		return false
	}

	if e.expectedValue != val {
		return false
	}

	e.actualCount++
	return true
}

func (e *expectation) isSatisfied() bool {
	if e.actualCount == e.expectedCount {
		return true
	}

	return false
}

