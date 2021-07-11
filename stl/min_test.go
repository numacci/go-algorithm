package stl

import "fmt"

func ExampleMin() {
	fmt.Println(min(-2, 3))
	fmt.Println(min(0, 8))

	// Output:
	// -2
	// 0
}

func ExampleMins() {
	fmt.Println(mins(-2, 3, 0, 1))
	fmt.Println(mins(8))

	// Output:
	// -2
	// 8
}

func ExampleChmin() {
	mi := 3
	fmt.Println(chmin(&mi, 10))
	fmt.Println(mi)

	mi = 3
	fmt.Println(chmin(&mi, -1))
	fmt.Println(mi)

	// Output:
	// false
	// 3
	// true
	// -1
}
