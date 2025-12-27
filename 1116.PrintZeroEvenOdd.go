package main

type ZeroEvenOdd struct {
	n           int
	chEvenReady chan struct{}
	chEvenDone  chan struct{}
	chOddReady  chan struct{}
	chOddDone   chan struct{}
}

func NewZeroEvenOdd(n int) *ZeroEvenOdd {
	zeo := &ZeroEvenOdd{
		n:           n,
		chEvenReady: make(chan struct{}),
		chEvenDone:  make(chan struct{}),
		chOddReady:  make(chan struct{}),
		chOddDone:   make(chan struct{}),
	}
	return zeo
}

func (z *ZeroEvenOdd) Zero(printNumber func(int)) {
	for i := 0; i < z.n; i++ {
		printNumber(0)
		if i%2 == 0 {
			z.chOddReady <- struct{}{}
			<-z.chOddDone
		} else {
			z.chEvenReady <- struct{}{}
			<-z.chEvenDone
		}
	}
}

func (z *ZeroEvenOdd) Even(printNumber func(int)) {
	for i := 2; i <= z.n; i += 2 {
		<-z.chEvenReady
		printNumber(i)
		z.chEvenDone <- struct{}{}
	}
}

func (z *ZeroEvenOdd) Odd(printNumber func(int)) {
	for i := 1; i <= z.n; i += 2 {
		<-z.chOddReady
		printNumber(i)
		z.chOddDone <- struct{}{}
	}
}

/*
func printn(a int) {
	fmt.Println(a)
}

func main() {
	z := NewZeroEvenOdd(4)
	go z.Odd(printn)
	go z.Even(printn)
	go z.Zero(printn)
	time.Sleep(time.Second)
}
*/
