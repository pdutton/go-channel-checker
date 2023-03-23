package checker

import (
    "fmt"
)

type ChannelChecker[T comparable] interface {
	Expect(T) Expectation[T]
	Check()
}

type Expectation[T comparable] interface {
    fmt.Stringer
    Times(int) Expectation[T]
    AtLeastOnce() Expectation[T]
    AnyTimes() Expectation[T]
    After(Expectation[T]) Expectation[T]
    matches(T) bool
    isSatisfied() bool
}

