package main

import (
	"fmt"
	"sort"
)

const MOD = 1000000007

func fun(diags []int, maxK int) []int64 {
	dp := make([]int64, maxK+1)
	dp[0] = 1

	for _, L := range diags {
		for j := maxK; j >= 1; j-- {
			ways := int64(L) - int64(j-1)
			if ways > 0 {
				dp[j] = (dp[j] + dp[j-1]*ways) % MOD
			}
		}
	}
	return dp
}

func task7() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		return
	}
	if k >= 2*n {
		fmt.Println(0)
		return
	}
	var whiteDiags []int
	var blackDiags []int

	for s := 0; s < 2*n-1; s++ {
		low := 0
		if s-(n-1) > 0 {
			low = s - (n - 1)
		}
		high := n - 1
		if s < n-1 {
			high = s
		}
		length := high - low + 1

		if s%2 == 0 {
			whiteDiags = append(whiteDiags, length)
		} else {
			blackDiags = append(blackDiags, length)
		}
	}

	sort.Ints(whiteDiags)
	sort.Ints(blackDiags)

	maxW := k
	if len(whiteDiags) < maxW {
		maxW = len(whiteDiags)
	}
	maxB := k
	if len(blackDiags) < maxB {
		maxB = len(blackDiags)
	}

	dpWhite := fun(whiteDiags, maxW)
	dpBlack := fun(blackDiags, maxB)

	var totalAns int64 = 0
	for i := 0; i <= maxW; i++ {
		j := k - i
		if j >= 0 && j <= maxB {
			contribution := (dpWhite[i] * dpBlack[j]) % MOD
			totalAns = (totalAns + contribution) % MOD
		}
	}

	fmt.Println(totalAns)
}

/*
func main() {
	task7()
}
*/
