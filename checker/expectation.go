package checker

import (
)

type expectation[T comparable] struct {
	expectedValue T
	expectedCount int
	actualCount int
}

func newExpectation[T comparable](val T, cnt int) *expectation[T] {
	return &expectation[T]{
		expectedValue: val,
		expectedCount: cnt,
	}
}

func (e *expectation[T]) matches(val T) bool {
	if e.actualCount >= e.expectedCount {
		return false
	}

	if e.expectedValue != val {
		return false
	}

	e.actualCount++
	return true
}

func (e *expectation[T]) isSatisfied() bool {
	if e.actualCount == e.expectedCount {
		return true
	}

	return false
}

