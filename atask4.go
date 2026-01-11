package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	next := func() int {
		scanner.Scan()
		val, _ := strconv.Atoi(scanner.Text())
		return val
	}

	n := next()
	k := next()

	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = next()
	}
	maxDeque := make([]int, 0)
	minDeque := make([]int, 0)

	var ans int64 = 0
	l := 0

	for r := 0; r < n; r++ {
		for len(maxDeque) > 0 && a[maxDeque[len(maxDeque)-1]] <= a[r] {
			maxDeque = maxDeque[:len(maxDeque)-1]
		}
		maxDeque = append(maxDeque, r)

		for len(minDeque) > 0 && a[minDeque[len(minDeque)-1]] >= a[r] {
			minDeque = minDeque[:len(minDeque)-1]
		}
		minDeque = append(minDeque, r)

		for a[maxDeque[0]]-a[minDeque[0]] > k {
			l++
			if maxDeque[0] < l {
				maxDeque = maxDeque[1:]
			}
			if minDeque[0] < l {
				minDeque = minDeque[1:]
			}
		}

		ans += int64(r - l + 1)
	}

	fmt.Println(ans)
}
