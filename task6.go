package main

import (
	"bufio"
	"os"
	"strconv"
)

type operation struct {
	l, r int64
}

func task6() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	next := func() string {
		if !scanner.Scan() {
			return ""
		}
		return scanner.Text()
	}
	nextInt64 := func() int64 {
		val, _ := strconv.ParseInt(next(), 10, 64)
		return val
	}
	nStr := next()
	if nStr == "" {
		return
	}
	qStr := next()
	q, _ := strconv.Atoi(qStr)
	sOrig := next()

	history := make([]operation, 0, q)

	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for i := 0; i < q; i++ {
		opType := nextInt64()

		if opType == 1 {
			l := nextInt64()
			r := nextInt64()
			history = append(history, operation{l: l, r: r})
		} else {
			pos := nextInt64()
			currPos := pos

			for j := len(history) - 1; j >= 0; j-- {
				hL := history[j].l
				hR := history[j].r
				subLen := hR - hL + 1
				newR := hR + subLen

				if currPos > newR {
					currPos -= subLen
				} else if currPos >= hL {
					currPos = hL + (currPos-hL)/2
				}
			}
			writer.WriteByte(sOrig[int(currPos-1)])
			writer.WriteByte('\n')
		}
	}
}

/*
func main() {
	task6()
}
*/
