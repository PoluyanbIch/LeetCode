package main

type Foo struct {
	ch1 chan struct{}
	ch2 chan struct{}
}

func NewFoo() *Foo {
	f := &Foo{
		ch1: make(chan struct{}),
		ch2: make(chan struct{}),
	}
	return f
}

func (f *Foo) First(printFirst func()) {
	// Do not change this line
	printFirst()
	f.ch1 <- struct{}{}
}

func (f *Foo) Second(printSecond func()) {
	/// Do not change this line
	<-f.ch1
	printSecond()
	f.ch2 <- struct{}{}
}

func (f *Foo) Third(printThird func()) {
	// Do not change this line
	<-f.ch2
	printThird()
}

/*
func main() {
	f := NewFoo()
	go f.Third(func() { fmt.Println("Third") })
	go f.First(func() { fmt.Println("First") })
	go f.Second(func() { fmt.Println("Second") })
	time.Sleep(time.Second)
}
*/
