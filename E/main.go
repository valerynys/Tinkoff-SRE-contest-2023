package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}

	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + a[i-1]
	}

	seen := make(map[int]bool)
	count := 0
	for i := 1; i <= n; i++ {
		for j := i; j <= n; j++ {
			s := prefix[j] - prefix[i-1]
			if j > i && seen[s] {
				count++
			}
			seen[s] = true
		}
	}

	fmt.Println(count + 1)
}
