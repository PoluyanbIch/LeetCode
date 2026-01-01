package main

func getCost(sub, target string) int {
	cost := 0
	for i := 0; i < 5; i++ {
		if sub[i] != target[i] {
			cost++
		}
	}
	return cost
}

func task2(s string) int {
	n := len(s)
	t1 := "tbank"
	t2 := "study"
	m := 5
	size := n - m + 1
	costT1 := make([]int, size)
	costT2 := make([]int, size)

	for i := 0; i < size; i++ {
		sub := s[i : i+m]

		costT1[i] = getCost(sub, t1)
		costT2[i] = getCost(sub, t2)
	}

	const INF = 10
	prefT1 := make([]int, size)
	prefT2 := make([]int, size)
	min1, min2 := INF, INF
	for i := 0; i < size; i++ {
		if costT1[i] < min1 {
			min1 = costT1[i]
		}
		if costT2[i] < min2 {
			min2 = costT2[i]
		}
		prefT1[i] = min1
		prefT2[i] = min2
	}

	ans := INF

	for i := m; i < size; i++ {
		if prefT1[i-m]+costT2[i] < ans {
			ans = prefT1[i-m] + costT2[i]
		}
		if prefT2[i-m]+costT1[i] < ans {
			ans = prefT2[i-m] + costT1[i]
		}
	}

	return ans
}

/*
func main() {
	var inp string
	fmt.Scan(&inp)
	fmt.Println(task2(inp))
}
*/
