package main

import "sync"

type DiningPhilosophers struct {
	forks [5]chan struct{}
	mu    sync.Mutex
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

func (this *DiningPhilosophers) wantsToEat(
	philosopher int,
	pickLeftFork func(),
	pickRightFork func(),
	eat func(),
	putLeftFork func(),
	putRightFork func(),
) {
	// TODO: implement your solution here
	this.mu.Lock()
	leftForkId := philosopher
	rightForkId := (philosopher + 1) % 5
	<-this.forks[leftForkId]
	pickLeftFork()
	<-this.forks[rightForkId]
	pickRightFork()
	eat()
	this.forks[leftForkId] <- struct{}{}
	putLeftFork()
	this.forks[rightForkId] <- struct{}{}
	putRightFork()
	this.mu.Unlock()
}
