package main

import (
	"math"
)

func task4(n int, adj [][]int) int {
	minCycle := math.MaxInt32
	for startNode := 0; startNode < n; startNode++ {
		dist := make([]int, n)
		parent := make([]int, n)
		for i := range dist {
			dist[i] = -1
			parent[i] = -1
		}

		queue := make([]int, 0, n)
		queue = append(queue, startNode)
		dist[startNode] = 0
		head := 0
		for head < len(queue) {
			u := queue[head]
			head++
			for _, v := range adj[u] {
				if dist[v] == -1 {
					dist[v] = dist[u] + 1
					parent[v] = u
					queue = append(queue, v)
				} else if parent[u] != v {
					currCycle := dist[u] + dist[v] + 1
					if currCycle < minCycle {
						minCycle = currCycle
					}
					if minCycle == 3 {
						return 3
					}
				}
			}
		}
	}
	if minCycle == math.MaxInt32 {
		return 0
	}
	return minCycle
}

/*
func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		return
	}
	adj := make([][]int, n)
	for i := 0; i < m; i++ {
		var u, v int
		if _, err := fmt.Scan(&u, &v); err != nil {
			break
		}
		u--
		v--
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	result := task4(n, adj)
	if result == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(result)
	}
}
*/
