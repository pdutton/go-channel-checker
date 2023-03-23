package checker

import (
    "strconv"
    "testing"
)

func TestNewExpectationTimes_string(t *testing.T) {
    var tests = []struct{
        val string
        nbr int
    }{
        { `red`, 1 },
        { `orange`, 3 },
    }

    for _, test := range tests {
        t.Run(test.val, func(t *testing.T) {
            t.Parallel()
            subtestTimes(t, test.val, test.nbr)
        })
    }
}

func TestNewExpectationTimes_int(t *testing.T) {
    var tests = []struct{
        val int
        nbr int
    }{
        { 9, 1 },
        { 3, 3 },
    }

    for _, test := range tests {
        t.Run(strconv.Itoa(test.val), func(t *testing.T) {
            t.Parallel()
            subtestTimes(t, test.val, test.nbr)
        })
    }
}

func subtestTimes[T comparable](t *testing.T, val T, nbr int) {
    var impl = newExpectation(val).Times(nbr).(*expectation[T])

    if impl.expectedValue != val {
        t.Errorf(`unexpected val: %v`, impl.expectedValue)
    }

    if impl.minCount != nbr {
        t.Errorf(`unexpected min count: %d`, impl.minCount)
    }

    if impl.maxCount != nbr {
        t.Errorf(`unexpected max count: %d`, impl.maxCount)
    }

    if impl.actualCount != 0 {
        t.Errorf(`unexpected actual count: %d`, impl.actualCount)
    }
}

func TestNewExpectationBetween(t *testing.T) {
    var tests = []struct{
        val string
        min int
        max int
    }{
        { `red`, 1, 2 },
        { `orange`, 2, 5 },
        { `yellow`, 0, 55 },
        { `green`, 4, 55 },
    }

    for _, test := range tests {
        t.Run(test.val, func(t *testing.T) {
            t.Parallel()
            subtestBetween(t, test.val, test.min, test.max)
        })
    }
}

func subtestBetween[T comparable](t *testing.T, val T, min int, max int) {
    var impl = newExpectation(val).Between(min, max).(*expectation[T])

    if impl.expectedValue != val {
        t.Errorf(`unexpected val: %v`, impl.expectedValue)
    }

    if impl.minCount != min {
        t.Errorf(`unexpected min count: %d`, impl.minCount)
    }

    if impl.maxCount != max {
        t.Errorf(`unexpected max count: %d`, impl.maxCount)
    }

    if impl.actualCount != 0 {
        t.Errorf(`unexpected actual count: %d`, impl.actualCount)
    }
}

func TestExpectationAtLeastOnce(t *testing.T) {
    var val = `asdf`
    var impl = newExpectation(val).AtLeastOnce().(*expectation[string])

    if impl.expectedValue != val {
        t.Errorf(`unexpected val: %v`, impl.expectedValue)
    }

    if impl.minCount != 1 {
        t.Errorf(`unexpected min count: %d`, impl.minCount)
    }

    if impl.maxCount != UNBOUNDED {
        t.Errorf(`unexpected max count: %d`, impl.maxCount)
    }

    if impl.actualCount != 0 {
        t.Errorf(`unexpected actual count: %d`, impl.actualCount)
    }
}

func TestExpectationAnyTimes(t *testing.T) {
    var val = `qwerty`
    var impl = newExpectation(val).AnyTimes().(*expectation[string])

    if impl.expectedValue != val {
        t.Errorf(`unexpected val: %v`, impl.expectedValue)
    }

    if impl.minCount != 0 {
        t.Errorf(`unexpected min count: %d`, impl.minCount)
    }

    if impl.maxCount != UNBOUNDED {
        t.Errorf(`unexpected max count: %d`, impl.maxCount)
    }

    if impl.actualCount != 0 {
        t.Errorf(`unexpected actual count: %d`, impl.actualCount)
    }
}

