package main

import "fmt"

func Checker(a, b int) int {
	if a < b {
		return -1
	}
	if a > b {
		return 1
	}
	if a == b {
		return 0
	}
	return 404
}

func IsSorted(f func(a, b int) int, a []int) bool {
	length := len(a)
	ascending := true
	descending := true

	for i := 0; i < length-1; i++ {
		cmp := f(a[i], a[i+1])

		if cmp > 0 {
			ascending = false
		} else if cmp < 0 {
			descending = false
		}

		if !(ascending || descending) {
			return false
		}
	}

	return true
}

func main() {
	a1 := []int{703794, 244404, 111349, -65949, -8, -628445, -635449, -822372}
	a2 := []int{0, 2, 1, 3}
	result1 := IsSorted(Checker, a1)
	result2 := IsSorted(Checker, a2)

	fmt.Println(result1)
	fmt.Println(result2)
}
