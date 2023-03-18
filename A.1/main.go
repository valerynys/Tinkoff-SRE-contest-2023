package main

import (
	"fmt"
	"sort"
)

func main() {
	var heights [4]int
	for i := 0; i < 4; i++ {
		fmt.Scan(&heights[i])
	}
	if sort.IntsAreSorted(heights[:]) || sort.IsSorted(sort.Reverse(sort.IntSlice(heights[:]))) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
