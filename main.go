package main

import "fmt"

func main() {
	grid := [][]int{
		{4, 3, 8, 4},
		{9, 5, 1, 9},
		{2, 7, 6, 2},
	}
	a := numMagicSquaresInside(grid)
	fmt.Println(a)
}
