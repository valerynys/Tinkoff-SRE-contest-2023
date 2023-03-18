package main

import (
	"fmt"
)

func main() {
	var n int
	var s string
	fmt.Scan(&n)
	fmt.Scan(&s)
	// словарь для хранения количества вхождений символов
	count := map[rune]int{'a': 0, 'b': 0, 'c': 0, 'd': 0}

	left, right := 0, 0           // два указателя
	min_len := int(^uint(0) >> 1) // минимальная длина хорошей подстроки

	for right < n {
		count[rune(s[right])]++
		right++

		// перемещаем левый указатель до тех пор, пока оставшиеся символы все еще встречаются
		for all(count) && left < right {
			count[rune(s[left])]--
			left++

			// обновляем минимальную длину хорошей подстроки
			if !all(count) {
				min_len = min(min_len, right-left+1)
			}
		}
	}

	if min_len == int(^uint(0)>>1) {
		fmt.Println(-1)
	} else {
		fmt.Println(min_len)
	}
}

// функция проверки, что все значения словаря больше 0
func all(count map[rune]int) bool {
	for _, v := range count {
		if v == 0 {
			return false
		}
	}
	return true
}

// функция для нахождения минимума двух чисел
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
