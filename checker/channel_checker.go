package checker

type ChannelChecker interface {
	Expect(string)
	Check()
}


