package checker

import (
	"context"
	"sync"
)

type channelCheckerImpl[T comparable] struct {
	t tIntf

	ch chan T

	expectationMutex sync.RWMutex
	expectations []*expectation[T]

	ctxt context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}

func NewChannelChecker[T comparable](t tIntf, ch chan T) ChannelChecker[T] {
	ctxt, cancel := context.WithCancel(context.Background())

	var ccp = &channelCheckerImpl[T]{
		t: t,

		ch: ch,

		ctxt: ctxt,
		cancel: cancel,
	}

	t.Cleanup(ccp.Check)

	go ccp.run()

	return ccp
}

func (cc *channelCheckerImpl[T]) run() {
	cc.wg.Add(1)
	defer cc.wg.Done()

	for {
		select {
		case <-cc.ctxt.Done():
			return
		case val, ok := <-cc.ch:
			if !ok {
				return
			}
			cc.verifyExpected(val)
		}
	}
}

func (cc *channelCheckerImpl[T]) verifyExpected(val T) {
	cc.expectationMutex.RLock()
	defer cc.expectationMutex.RUnlock()

	for _, e := range cc.expectations {
		if e.matches(val) {
			return
		}
	}

	cc.t.Errorf(`unexpected value received on channel: %s`, val)

}

func (cc *channelCheckerImpl[T]) Expect(val T) {
	cc.expectationMutex.Lock()
	defer cc.expectationMutex.Unlock()

	var expect = newExpectation(val, 1)
	cc.expectations = append(cc.expectations, expect)

}

func (cc *channelCheckerImpl[T]) Check() {
	cc.cancel()
	cc.wg.Wait()

	cc.expectationMutex.RLock()
	defer cc.expectationMutex.RUnlock()

	for _, e := range cc.expectations {
		if !e.isSatisfied() {
			cc.t.Errorf(`expected to receive %s on channel`, e.expectedValue)
		}
	}

}

