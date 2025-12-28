package main

import (
	"fmt"
	"time"
)

func main() {
	dp := NewDiningPhilosophers()
	go dp.wantsToEat(2,
		func() { fmt.Println("2 pick left fork") },
		func() { fmt.Println("2 pick right fork") },
		func() { fmt.Println("2 eat") },
		func() { fmt.Println("2 put left fork") },
		func() { fmt.Println("2 put right fork") },
	)
	go dp.wantsToEat(1,
		func() { fmt.Println("1 pick left fork") },
		func() { fmt.Println("1 pick right fork") },
		func() { fmt.Println("1 eat") },
		func() { fmt.Println("1 put left fork") },
		func() { fmt.Println("1 put right fork") },
	)
	go dp.wantsToEat(3,
		func() { fmt.Println("3 pick left fork") },
		func() { fmt.Println("3 pick right fork") },
		func() { fmt.Println("3 eat") },
		func() { fmt.Println("3 put left fork") },
		func() { fmt.Println("3 put right fork") },
	)
	go dp.wantsToEat(4,
		func() { fmt.Println("4 pick left fork") },
		func() { fmt.Println("4 pick right fork") },
		func() { fmt.Println("4 eat") },
		func() { fmt.Println("4 put left fork") },
		func() { fmt.Println("4 put right fork") },
	)
	go dp.wantsToEat(0,
		func() { fmt.Println("0 pick left fork") },
		func() { fmt.Println("0 pick right fork") },
		func() { fmt.Println("0 eat") },
		func() { fmt.Println("0 put left fork") },
		func() { fmt.Println("0 put right fork") },
	)
	time.Sleep(time.Second * 5)
}
