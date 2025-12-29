package main

type DiningPhilosophers struct {
	forks [5]chan struct{}
}

func NewDiningPhilosophers() *DiningPhilosophers {
	var f [5]chan struct{}
	dp := &DiningPhilosophers{forks: f}
	for i := 0; i < 5; i++ {
		dp.forks[i] = make(chan struct{}, 1)
		dp.forks[i] <- struct{}{}
	}
	return dp
}

func (dp *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	left := philosopher
	right := (philosopher + 1) % 5

	if philosopher == 4 {
		<-dp.forks[right]
		pickRightFork()
		<-dp.forks[left]
		pickLeftFork()
	} else {
		<-dp.forks[left]
		pickLeftFork()
		<-dp.forks[right]
		pickRightFork()
	}

	eat()

	dp.forks[left] <- struct{}{}
	dp.forks[right] <- struct{}{}
	putLeftFork()
	putRightFork()
}
