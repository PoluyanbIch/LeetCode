package main

func numMagicSquaresInside(grid [][]int) int {
	res := 0
	for a := 0; a <= len(grid)-3; a++ {
		for b := 0; b <= len(grid[0])-3; b++ {
			s := [10]bool{}
			isMagic := true
			for i := 0; i < 3 && isMagic; i++ {
				for j := 0; j < 3 && isMagic; j++ {
					if grid[a+i][b+j] < 1 || grid[a+i][b+j] > 9 {
						isMagic = false
					}
					if isMagic && s[grid[a+i][b+j]] {
						isMagic = false
					}
					if isMagic {
						s[grid[a+i][b+j]] = true
					}
				}
			}
			for _, el := range s[1:] {
				if !el {
					isMagic = false
					break
				}
			}
			if !isMagic {
				continue
			}
			row1 := grid[a][b] + grid[a][b+1] + grid[a][b+2]
			row2 := grid[a+1][b] + grid[a+1][b+1] + grid[a+1][b+2]
			row3 := grid[a+2][b] + grid[a+2][b+1] + grid[a+2][b+2]
			col1 := grid[a][b] + grid[a+1][b] + grid[a+2][b]
			col2 := grid[a][b+1] + grid[a+1][b+1] + grid[a+2][b+1]
			col3 := grid[a][b+2] + grid[a+1][b+2] + grid[a+2][b+2]
			diag1 := grid[a][b] + grid[a+1][b+1] + grid[a+2][b+2]
			diag2 := grid[a][b+2] + grid[a+1][b+1] + grid[a+2][b]
			m := [8]int{row1, row2, row3, col1, col2, col3, diag1, diag2}
			allEqual := true
			for i := 1; i < len(m); i++ {
				if m[i-1] != m[i] {
					allEqual = false
					break
				}
			}
			if allEqual {
				res++
			}
		}
	}
	return res
}
