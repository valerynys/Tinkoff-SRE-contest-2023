package main

import (
	"fmt"
)

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	time := ((n*k + m - 1) / m) * 1
	fmt.Println(time)
}
