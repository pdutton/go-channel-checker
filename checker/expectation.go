package checker

import (
    "fmt"
)

const UNBOUNDED = -1

type expectation[T comparable] struct {
	expectedValue T
    minCount int
    maxCount int
	actualCount int
    after Expectation[T]
}

func newExpectation[T comparable](val T) Expectation[T] {
	return &expectation[T]{
		expectedValue: val,
        minCount: 1,
        maxCount: 1,
	}
}

func (e *expectation[T]) Times(nbr int) Expectation[T] {
    e.minCount = nbr
    e.maxCount = nbr

    return e
}

func (e *expectation[T]) Between(min, max int) Expectation[T] {
    e.minCount = min
    e.maxCount = max

    return e
}

func (e *expectation[T]) AtLeastOnce() Expectation[T] {
    e.minCount = 1
    e.maxCount = UNBOUNDED

    return e
}

func (e *expectation[T]) AnyTimes() Expectation[T] {
    e.minCount = 0
    e.maxCount = UNBOUNDED

    return e
}


func (e *expectation[T]) After(other Expectation[T]) Expectation[T] {
    e.after = other
    return e
}

func (e *expectation[T]) matches(val T) bool {
    if e.after != nil && !e.after.isSatisfied() {
        return false
    }
    
	if e.maxCount != UNBOUNDED && 
       e.maxCount <= e.actualCount {
		return false
	}

	if e.expectedValue != val {
		return false
	}

	e.actualCount++

	return true
}

func (e *expectation[T]) isSatisfied() bool {

    if e.minCount == 0 {
        return true
    }

    return (e.actualCount >= e.minCount)
}

func (e *expectation[T]) String() string {
    return fmt.Sprintf(`expected to receive %v`, e.expectedValue)
}

