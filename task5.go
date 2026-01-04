package main

import (
	"fmt"
	"strconv"
	"strings"
)

func task5() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		return
	}
	a := make([]int, n)
	cnt := make(map[int]int)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
		cnt[a[i]]++
	}
	res := make([]string, n)
	for i := 0; i < n; i++ {
		x := a[i]
		var ans int
		if cnt[x] >= 2 {
			ans = n - cnt[x]
		} else {
			left := a[(i-1+n)%n]
			right := a[(i+1)%n]
			if (left >= x && right >= x) || (left <= x && right <= x) {
				ans = n - 1
			} else {
				ans = n
			}
		}
		res[i] = strconv.Itoa(ans)
	}
	fmt.Println(strings.Join(res, " "))
}

/*
func main() {
	task5()
}
*/
