package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	var heights [4]int
	for i := 0; i < 4; i++ {
		scanner.Scan()
		heights[i], _ = strconv.Atoi(scanner.Text())
	}

	if heights[0] < heights[1] && heights[1] < heights[2] && heights[2] < heights[3] {
		fmt.Println("YES")
	} else if heights[0] > heights[1] && heights[1] > heights[2] && heights[2] > heights[3] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
