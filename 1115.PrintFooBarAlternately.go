package main

type FooBar struct {
	n   int
	chF chan struct{}
	chB chan struct{}
}

func NewFooBar(n int) *FooBar {
	return &FooBar{
		n:   n,
		chF: make(chan struct{}),
		chB: make(chan struct{}),
	}
}

func (fb *FooBar) Foo(printFoo func()) {
	for i := 0; i < fb.n; i++ {
		// printFoo() outputs "foo". Do not change or remove this line.
		printFoo()
		fb.chF <- struct{}{}
		<-fb.chB
	}
}

func (fb *FooBar) Bar(printBar func()) {
	for i := 0; i < fb.n; i++ {
		// printBar() outputs "bar". Do not change or remove this line.
		<-fb.chF
		printBar()
		fb.chB <- struct{}{}
	}
}

/*
func main() {
	f := NewFooBar(3)
	go f.Foo(func() { fmt.Println("Foo") })
	time.Sleep(time.Second)
	go f.Bar(func() { fmt.Println("Bar") })
	time.Sleep(time.Second)
}
*/
