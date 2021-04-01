package findsumpair

// O(2) quadratic
func HasSumPairOrder2(ascOrderNums []int, sum int) (found bool, index1 int, index2 int) {
	for i, c1 := range ascOrderNums {
		for j, c2 := range ascOrderNums {
			if i == j {
				continue
			}
			if c1+c2 == sum {
				return true, i, j
			}
		}
	}
	return false, -1, -1
}

// O(n) linear but requires ordered array
func HasSumPairOrder1Ordered(ascOrderNums []int, sum int) (found bool, index1 int, index2 int) {
	i, j := 0, len(ascOrderNums)-1
	for {
		if i >= j {
			break
		}
		proposal := ascOrderNums[i] + ascOrderNums[j]
		if proposal == sum {
			return true, i, j
		}
		if proposal > sum {
			j--
			continue
		}
		i++
	}
	return false, -1, -1
}

// O(n) linear but requires constant time hash set lookup
func HasSumPairOrder1Unordered(unorderedNums []int, sum int) (found bool, index1 int, index2 int) {
	seen := make(map[int]int)
	for i, c1 := range unorderedNums {
		neededComp := sum - c1
		if pos, yes := seen[neededComp]; yes {
			return true, pos, i
		}
		seen[c1] = i
	}
	return false, -1, -1
}
