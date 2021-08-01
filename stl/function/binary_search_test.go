package function

import (
	"fmt"
	"sort"
)

func ExampleLowerBound() {
	// target slice should be sorted
	items := []int{1, 3, 4, 4, 6, 9}
	fmt.Println(sort.Search(len(items), func(i int) bool {
		return items[i] >= 4
	}))
	fmt.Println(sort.Search(len(items), func(i int) bool {
		return items[i] >= 6
	}))

	// Output:
	// 2
	// 4
}

func ExampleUpperBound() {
	// target slice should be sorted
	items := []float64{1.0, 3.2, 4.8, 4.8, 6.5, 9.0}
	fmt.Println(sort.Search(len(items), func(i int) bool {
		return items[i] > 2.8
	}))
	fmt.Println(sort.Search(len(items), func(i int) bool {
		return items[i] > 4.8
	}))

	// Output:
	// 1
	// 4
}
