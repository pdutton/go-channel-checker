package checker

import (
    "fmt"
)

const (
    ONE_OR_MORE_TIMES  = -1
    ZERO_OR_MORE_TIMES = -2
)

type expectation[T comparable] struct {
	expectedValue T
	expectedCount int
	actualCount int
    after Expectation[T]
}

func newExpectation[T comparable](val T) Expectation[T] {
	return &expectation[T]{
		expectedValue: val,
		expectedCount: 1,
	}
}

func (e *expectation[T]) Times(nbr int) Expectation[T] {
    e.expectedCount = nbr
    return e
}

func (e *expectation[T]) AtLeastOnce() Expectation[T] {
    e.expectedCount = ONE_OR_MORE_TIMES
    return e
}

func (e *expectation[T]) AnyTimes() Expectation[T] {
    e.expectedCount = ZERO_OR_MORE_TIMES
    return e
}


func (e *expectation[T]) After(other Expectation[T]) Expectation[T] {
    e.after = other
    return e
}

func (e *expectation[T]) matches(val T) bool {
	if e.expectedCount >= 0 && 
       e.expectedCount <= e.actualCount {
		return false
	}

	if e.expectedValue != val {
		return false
	}

	e.actualCount++

	return true
}

func (e *expectation[T]) isSatisfied() bool {

    if e.expectedCount == ZERO_OR_MORE_TIMES {
        return true
    }

    if e.expectedCount == ONE_OR_MORE_TIMES {
        return (e.actualCount > 0) 
    }

	if e.actualCount == e.expectedCount {
		return true
	}

	return false
}

func (e *expectation[T]) String() string {
    return fmt.Sprintf(`expected to receive %v`, e.expectedValue)
}

