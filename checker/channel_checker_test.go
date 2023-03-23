package checker

import (
	"github.com/golang/mock/gomock"
	"github.com/pdutton/go-channel-checker/mock"
	"testing"
)

func TestChannelChecker_noExpectations_noSends(t *testing.T) {
	var mc = gomock.MockController(t)
	var mt = mock.NewMockTB(mc)

	var ch = make(chan string)
	var _ = NewChannelChecker(t, ch)
}

func TestChannelChecker_noExpectations_oneSends(t *testing.T) {
	var ch = make(chan string)
	var _ = NewChannelChecker(t, ch)

	ch<-`hi`
}

func TestChannelChecker_oneExpectation_noSends(t *testing.T) {
	var ch = make(chan string)
	var cc = NewChannelChecker(t, ch)

	cc.Expect(`hi`)
}

func TestChannelChecker_oneExpectation_oneSend(t *testing.T) {
	var ch = make(chan string)
	var cc = NewChannelChecker(t, ch)

	cc.Expect(`hi`)

	ch<-`hi`
}

