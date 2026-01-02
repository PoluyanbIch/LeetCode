package main

import (
	"fmt"
)

func task3() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		return
	}
	for i := 0; i < t; i++ {
		var s string
		fmt.Scan(&s)
		n := int64(len(s))
		hasZero := false
		for j := 0; j < len(s); j++ {
			if s[j] == '0' {
				hasZero = true
				break
			}
		}
		if !hasZero {
			fmt.Println(n * n)
			continue
		}
		maxInner := 0
		cur := 0
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				cur++
				if cur > maxInner {
					maxInner = cur
				}
			} else {
				cur = 0
			}
		}
		pref := 0
		for j := 0; j < len(s); j++ {
			if s[j] == '1' {
				pref++
			} else {
				break
			}
		}
		suff := 0
		for j := len(s) - 1; j >= 0; j-- {
			if s[j] == '1' {
				suff++
			} else {
				break
			}
		}
		m := int64(maxInner)
		if int64(pref+suff) > m {
			m = int64(pref + suff)
		}
		ans := (m + 1) * (m + 1) / 4
		fmt.Println(ans)
	}
}

/*
func main() {
	task3()
}
*/
