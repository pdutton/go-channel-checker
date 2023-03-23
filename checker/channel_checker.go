package checker

type ChannelChecker[T any] interface {
	Expect(T)
	Check()
}


