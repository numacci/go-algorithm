package stl

import (
	"fmt"
	"sort"
)

func ExampleSort() {
	// case int
	intS := []int{-3, 2, 0, 8, -5, 1}
	sort.Ints(intS)
	fmt.Println(intS)

	// case float64
	floatS := []float64{-3.5, 2.5, 0.5, 8.5, -5.5, 1.5}
	sort.Float64s(floatS)
	fmt.Println(floatS)

	// case string
	stringS := []string{"orange", "lemon", "banana", "apple"}
	sort.Strings(stringS)
	fmt.Println(stringS)

	// case struct
	structS := []struct {
		to   int
		cost int
	}{
		{3, 1},
		{2, 10},
		{4, 4},
	}
	sort.Slice(structS, func(i, j int) bool {
		// define Less function
		return structS[i].cost < structS[j].cost
	})
	fmt.Println(structS)

	// Output:
	// [-5 -3 0 1 2 8]
	// [-5.5 -3.5 0.5 1.5 2.5 8.5]
	// [apple banana lemon orange]
	// [{3 1} {4 4} {2 10}]
}
