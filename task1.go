package main

import (
	"sort"
	"strings"
)

func task1(inputStr string) string {
	s := strings.ReplaceAll(inputStr, "0", "")
	if len(s) == 0 {
		return "0"
	}
	chars := []rune(inputStr)
	sort.Slice(chars, func(i, j int) bool {
		return chars[i] < chars[j]
	})
	if chars[0] == '0' {
		for i := 0; i < len(chars); i++ {
			if chars[i] != '0' {
				chars[0], chars[i] = chars[i], chars[0]
				break
			}
		}
	}
	result := strings.TrimLeft(string(chars), "0")
	return result
}

/*
func main() {
	var inp string
	fmt.Scan(&inp)
	fmt.Println(task1(inp))
}
*/
