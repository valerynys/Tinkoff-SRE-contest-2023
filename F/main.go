package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, s int
	fmt.Scan(&n, &s)

	students := make([][2]int, n)
	for i := 0; i < n; i++ {
		var l, r int
		fmt.Scan(&l, &r)
		students[i] = [2]int{l, r}
	}

	sort.Slice(students, func(i, j int) bool {
		return students[i][1] < students[j][1]
	})

	lo, hi := students[(n-1)/2][0], students[(n-1)/2][1]
	for lo < hi {
		mid := (lo + hi + 1) / 2
		leftSum, rightSum := 0, 0
		for i, s := range students {
			if i <= (n-1)/2 {
				leftSum += max(s[0], mid)
			} else {
				rightSum += max(s[0], mid)
			}
		}
		if leftSum+rightSum > s {
			hi = mid - 1
		} else {
			lo = mid
		}
	}

	fmt.Println(lo + 1)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
