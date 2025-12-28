package main

import (
	"fmt"
	"sync/atomic"
)

type H2O struct {
	chH chan struct{}
	chO chan struct{}
	a   atomic.Int64
}

func NewH2O() *H2O {
	h := &H2O{
		chH: make(chan struct{}, 2),
		chO: make(chan struct{}, 2),
	}
	h.chO <- struct{}{}
	h.chO <- struct{}{}
	return h
}

func (h *H2O) Hydrogen(releaseHydrogen func()) {
	// releaseHydrogen() outputs "H". Do not change or remove this line.
	<-h.chO
	releaseHydrogen()
	h.a.Add(1)
	if h.a.CompareAndSwap(2, 0) {
		h.chH <- struct{}{}
	}
}

func (h *H2O) Oxygen(releaseOxygen func()) {
	// releaseOxygen() outputs "O". Do not change or remove this line.
	<-h.chH
	releaseOxygen()
	h.chO <- struct{}{}
	h.chO <- struct{}{}
}

func printH() {
	fmt.Println("H")
}

func printO() {
	fmt.Println("O")
}
