package function

import "fmt"

func ExampleLowerBound() {
	// target slice should be sorted
	items := []int{1, 3, 4, 4, 6, 9}
	fmt.Println(LowerBound(items, 4))
	fmt.Println(LowerBound(items, 6))

	// Output:
	// 2
	// 4
}

func ExampleUpperBound() {
	// target slice should be sorted
	items := []float64{1.0, 3.2, 4.8, 4.8, 6.5, 9.0}
	fmt.Println(UpperBound(items, 2.8))
	fmt.Println(UpperBound(items, 4.8))

	// Output:
	// 1
	// 4
}
