package main

type FenwickTree struct {
	tree []int
}

func newFenwickTree(n int) *FenwickTree {
	return &FenwickTree{tree: make([]int, n+1)}
}

func (f *FenwickTree) update(i int, val int) {
	for i < len(f.tree) {
		f.tree[i] += val
		i += i & -i
	}
}

func (f *FenwickTree) query(i int) int {
	sum := 0
	for i > 0 {
		sum += f.tree[i]
		i -= i & -i
	}
	return sum
}

func (f *FenwickTree) queryRange(l, r int) int {
	if l > r {
		return 0
	}
	return f.query(r) - f.query(l-1)
}

/*
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	next := func() string {
		scanner.Scan()
		return scanner.Text()
	}

	nextInt := func() int {
		v, _ := strconv.Atoi(next())
		return v
	}

	nextInt64 := func() int64 {
		v, _ := strconv.ParseInt(next(), 10, 64)
		return v
	}

	n := nextInt()
	k := nextInt64()

	a := make([]int64, n)
	for i := 0; i < n; i++ {
		a[i] = nextInt64()
	}

	sortedUnique := make([]int64, n)
	copy(sortedUnique, a)
	sort.Slice(sortedUnique, func(i, j int) bool { return sortedUnique[i] < sortedUnique[j] })

	m := 0
	for i := 0; i < n; i++ {
		if i == 0 || sortedUnique[i] != sortedUnique[i-1] {
			sortedUnique[m] = sortedUnique[i]
			m++
		}
	}
	sortedUnique = sortedUnique[:m]

	ft := newFenwickTree(m)
	isReachedHeight := make([]bool, m)

	results := make([]int, n)

	for i := 0; i < n; i++ {
		h := a[i]

		currIdx := sort.Search(len(sortedUnique), func(j int) bool {
			return sortedUnique[j] >= h
		})

		canReach := false
		if i == 0 {
			canReach = true
		} else {
			L := sort.Search(len(sortedUnique), func(j int) bool {
				return sortedUnique[j] >= h-k
			})
			R := sort.Search(len(sortedUnique), func(j int) bool {
				return sortedUnique[j] > h+k
			}) - 1

			if ft.queryRange(L+1, R+1) > 0 {
				canReach = true
			}
		}

		if canReach {
			results[i] = 1
			if !isReachedHeight[currIdx] {
				ft.update(currIdx+1, 1)
				isReachedHeight[currIdx] = true
			}
		} else {
			results[i] = 0
		}
	}
	for i := 0; i < n; i++ {
		writer.WriteString(strconv.Itoa(results[i]))
		if i < n-1 {
			writer.WriteByte(' ')
		}
	}
	writer.WriteByte('\n')
}
*/
