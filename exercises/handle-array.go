package exercises

import "fmt"

func handleArray[K, V any](arr []K, f func(K) V) []V {

	result := make([]V, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func HandleArray() {

	arr1 := []int{1, 2, 3}

	rs := handleArray[int](arr1, func(v int) int {
		return v * 2
	})

	fmt.Println(rs)
}
