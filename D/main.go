package main

import "fmt"

func isBoring(prefix []int) bool {
	d := make(map[int]int)
	for _, x := range prefix {
		d[x]++
	}
	var values []int
	for _, v := range d {
		values = append(values, v)
	}
	for i := 0; i < len(values); i++ {
		if i == 0 && len(values) == 1 {
			continue
		}
		if i == 0 && values[i] == values[i+1] {
			continue
		}
		if i == len(values)-1 && values[i-1] == values[i] {
			continue
		}
		newValues := append(values[:i], values[i+1:]...)
		if len(newValues) == len(values)-1 && allEqual(newValues) {
			return true
		}
	}
	return false
}

func allEqual(values []int) bool {
	if len(values) == 0 {
		return true
	}
	first := values[0]
	for i := 1; i < len(values); i++ {
		if values[i] != first {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	for l := n; l > 1; l-- {
		if isBoring(a[:l]) {
			fmt.Println(l)
			return
		}
	}
	fmt.Println(2)
}
