package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func readInt(r *bufio.Reader) (int, error) {
	res := 0
	found := false
	for {
		b, err := r.ReadByte()
		if err != nil {
			if found {
				return res, nil
			}
			return 0, err
		}
		if b >= '0' && b <= '9' {
			res = res*10 + int(b-'0')
			found = true
		} else if found {
			return res, nil
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	n, err := readInt(reader)
	if err != nil {
		return
	}

	n1, n5 := 0, 0

	for i := 0; i < n; i++ {
		r, err := readInt(reader)
		if err != nil {
			break
		}

		k := int(float64(r) / math.Sqrt(2.0))

		rSq := int64(r) * int64(r)
		for int64(k+1)*int64(k+1)*2 <= rSq {
			k++
		}
		for k > 0 && int64(k)*int64(k)*2 > rSq {
			k--
		}

		if (r+k)%2 == 0 {
			n1++
		} else {
			n5++
		}
	}
	currentSumMod8 := (n1*1 + n5*5) % 8

	ans := n
	for x := 0; x <= 8 && x <= n1; x++ {
		for y := 0; y <= 8 && y <= n5; y++ {
			if (x*1+y*5)%8 == currentSumMod8 {
				if x+y < ans {
					ans = x + y
				}
			}
		}
	}

	fmt.Println(ans)
}
