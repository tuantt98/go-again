package basic

import "fmt"

func mapArray[K comparable](arr []K, f func(K) K) []K {
	result := make([]K, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func Basic6() {
	a := []int{1, 2, 3}

	result := mapArray[int](a, func(i int) int {
		return i * 2
	})

	fmt.Println("Result: ", result)
}
