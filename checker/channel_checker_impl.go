package checker

import (
	"context"
	"fmt"
	"sync"
)

type channelCheckerImpl struct {
	t T	

	ch chan string	

	expectationMutex sync.RWMutex
	expectations []*expectation

	ctxt context.Context
	cancel context.CancelFunc
	wg sync.WaitGroup
}

func NewChannelChecker(t T, ch chan string) ChannelChecker {
	ctxt, cancel := context.WithCancel(context.Background())

	var ccp = &channelCheckerImpl{
		t: t,

		ch: ch,

		ctxt: ctxt,
		cancel: cancel,
	}

	t.Cleanup(ccp.Check)

	go ccp.run()

	return ccp
}

func (cc *channelCheckerImpl) run() {
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
			fmt.Println(`receive`)
			cc.verifyExpected(val)
		}
	}
}

func (cc *channelCheckerImpl) verifyExpected(val string) {
	cc.expectationMutex.RLock()
	defer cc.expectationMutex.RUnlock()

	fmt.Println(`verifying`)
	for _, e := range cc.expectations {
		if e.matches(val) {
			fmt.Println(`match found`)
			return
		}
	}

	cc.t.Errorf(`unexpected value received on channel: %s`, val)

}

func (cc *channelCheckerImpl) Expect(val string) {
	cc.expectationMutex.Lock()
	defer cc.expectationMutex.Unlock()

	var expect = newExpectation(val, 1)
	cc.expectations = append(cc.expectations, expect)

}

func (cc *channelCheckerImpl) Check() {
	cc.cancel()
	cc.wg.Wait()

	cc.expectationMutex.RLock()
	defer cc.expectationMutex.RUnlock()

	fmt.Println(`checking`)

	for _, e := range cc.expectations {
		if !e.isSatisfied() {
			cc.t.Errorf(`expected to receive %s on channel`, e.expectedValue)
		}
	}

}

