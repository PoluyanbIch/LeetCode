package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	next := func() string {
		if scanner.Scan() {
			return scanner.Text()
		}
		return ""
	}

	nStr := next()
	if nStr == "" {
		return
	}
	n, _ := strconv.Atoi(nStr)

	kStr := next()
	if kStr == "" {
		return
	}
	k, _ := strconv.Atoi(kStr)

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		val, _ := strconv.ParseInt(next(), 10, 64)
		a[i] = val
	}
	prefixSum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		prefixSum[i] = prefixSum[i-1] + a[i-1]
	}

	v := make([]int64, 0, n)
	for i := 2; i < n; i++ {
		v = append(v, prefixSum[i])
	}

	sort.Slice(v, func(i, j int) bool {
		return v[i] > v[j]
	})

	var maxScore int64 = 0

	var scoreNotEnding int64 = 0
	for i := 0; i < len(v) && i < k; i++ {
		if v[i] > 0 {
			scoreNotEnding += v[i]
		} else {
			break
		}
	}
	if scoreNotEnding > maxScore {
		maxScore = scoreNotEnding
	}

	var scoreEnding int64 = prefixSum[n]
	for i := 0; i < len(v) && i < k-1; i++ {
		if v[i] > 0 {
			scoreEnding += v[i]
		} else {
			break
		}
	}
	if scoreEnding > maxScore {
		maxScore = scoreEnding
	}

	fmt.Println(maxScore)
}
