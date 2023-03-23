package checker

import (
	"github.com/golang/mock/gomock"
	"testing"
)

func TestChannelChecker_noExpectations_noSends(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)

	mt.EXPECT().Cleanup(gomock.Any())

	var cc = NewChannelChecker(mt, ch)

	cc.Check()
}

func TestChannelChecker_noExpectations_oneSends(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var val = `hi`

	mt.EXPECT().Cleanup(gomock.Any())
	mt.EXPECT().Errorf(gomock.Any(), val)

	var cc = NewChannelChecker(mt, ch)

	ch<-val

	cc.Check()
}

func TestChannelChecker_oneExpectation_noSends(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var val = `hi`

	mt.EXPECT().Cleanup(gomock.Any())
	mt.EXPECT().Errorf(gomock.Any(), val)

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(val)

	cc.Check()
}

func TestChannelChecker_dupSend(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var val = `hi`

	mt.EXPECT().Cleanup(gomock.Any())
	mt.EXPECT().Errorf(gomock.Any(), val)

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(val)

	ch<-val
	ch<-val

	cc.Check()
}

func TestChannelChecker_dupExpect(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var val = `hi`

	mt.EXPECT().Cleanup(gomock.Any())
	mt.EXPECT().Errorf(gomock.Any(), val)

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(val)
	cc.Expect(val)

	ch<-val

	cc.Check()
}

func TestChannelChecker_oneExpectation_oneSend(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var val = `hi`

	mt.EXPECT().Cleanup(gomock.Any())

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(val)

	ch<-val

	cc.Check()
}

func TestChannelChecker_many_success_ordered(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var values = []string{`alpha`, `beta`, `delta`, `gamma`, `eplison`}

	mt.EXPECT().Cleanup(gomock.Any())

	var cc = NewChannelChecker(mt, ch)

	for _, v := range values {
		cc.Expect(v)
	}

	for _, v := range values {
		ch<-v
	}

	cc.Check()
}

func TestChannelChecker_many_success_disordered(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var values = []string{`alpha`, `beta`, `delta`, `gamma`, `eplison`}

	mt.EXPECT().Cleanup(gomock.Any())

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(values[0])
	cc.Expect(values[1])
	cc.Expect(values[2])
	cc.Expect(values[3])
	cc.Expect(values[4])

	ch<-values[2]
	ch<-values[1]
	ch<-values[0]
	ch<-values[4]
	ch<-values[3]

	cc.Check()
}

func TestChannelChecker_many_someProblems(t *testing.T) {
	var mc = gomock.NewController(t)
	var mt = NewMocktIntf(mc)
	var ch = make(chan string)
	var values = []string{`alpha`, `beta`, `delta`, `gamma`, `eplison`}

	mt.EXPECT().Cleanup(gomock.Any())
	mt.EXPECT().Errorf(gomock.Any(), values[0])
	mt.EXPECT().Errorf(gomock.Any(), values[4])

	var cc = NewChannelChecker(mt, ch)

	cc.Expect(values[0])
	cc.Expect(values[1])
	cc.Expect(values[2])
	cc.Expect(values[3])

	ch<-values[2]
	ch<-values[3]
	ch<-values[4]
	ch<-values[1]

	cc.Check()
}


