package function

import "fmt"

func ExampleMin() {
	fmt.Println(Min(-2, 3))
	fmt.Println(Min(0, 8))

	// Output:
	// -2
	// 0
}

func ExampleMins() {
	fmt.Println(Mins(-2, 3, 0, 1))
	fmt.Println(Mins(8))

	// Output:
	// -2
	// 8
}

func ExampleChmin() {
	mi := 3
	fmt.Println(Chmin(&mi, 10))
	fmt.Println(mi)

	mi = 3
	fmt.Println(Chmin(&mi, -1))
	fmt.Println(mi)

	// Output:
	// false
	// 3
	// true
	// -1
}
