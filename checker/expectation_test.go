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
        { `yellow`, ONE_OR_MORE_TIMES },
        { `green`, ZERO_OR_MORE_TIMES },
    }

    for _, test := range tests {
        t.Run(test.val, func(t *testing.T) {
            t.Parallel()
            subtest(t, test.val, test.nbr)
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
        { 4, ONE_OR_MORE_TIMES },
        { 6, ZERO_OR_MORE_TIMES },
    }

    for _, test := range tests {
        t.Run(strconv.Itoa(test.val), func(t *testing.T) {
            t.Parallel()
            subtest(t, test.val, test.nbr)
        })
    }
}

func subtest[T comparable](t *testing.T, val T, nbr int) {
    var impl = newExpectation(val).Times(nbr).(*expectation[T])

    if impl.expectedValue != val {
        t.Errorf(`unexpected val: %v`, impl.expectedValue)
    }

    if impl.expectedCount != nbr {
        t.Errorf(`unexpected count: %d`, impl.expectedCount)
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

    if impl.expectedCount != ONE_OR_MORE_TIMES {
        t.Errorf(`unexpected count: %d`, impl.expectedCount)
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

    if impl.expectedCount != ZERO_OR_MORE_TIMES {
        t.Errorf(`unexpected count: %d`, impl.expectedCount)
    }

    if impl.actualCount != 0 {
        t.Errorf(`unexpected actual count: %d`, impl.actualCount)
    }
}


