package basic

import (
	"fmt"
	"strings"
)

func mapAny[K, V any](arr []K, f func(K) V) []V {
	result := make([]V, len(arr))

	for i, v := range arr {
		result[i] = f(v)
	}

	return result
}

func Basic7() {
	a := []int{1, 2, 3, 4, 5, 6, 7}

	result := mapAny[int](a, func(i int) int {
		return i * 2
	})
	fmt.Println("Result: ", result)

	b := []string{"apply", "samsung"}

	resultStr := mapAny[string](b, func(i string) string {
		return strings.ToUpper(i)
	})

	fmt.Println(resultStr)

	c := make([]int, 0, 5)
	d := []int{1, 2, 3}
	c = append(c, d...)

	fmt.Println(c, len(c), cap(c))
}
