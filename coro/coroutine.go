package coro

func New[In, Out any](f func(In, func(Out) In) Out) (resume func(In) Out) {
	cin := make(chan In)
	cout := make(chan Out)
	resume = func(in In) Out {
		cin <- in
		return <-cout
	}
	yield := func(out Out) In {
		cout <- out
		return <-cin
	}
	go func() {
		cout <- f(<-cin, yield)
	}()
	return resume
}