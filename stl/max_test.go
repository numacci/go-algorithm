package stl

import "fmt"

func ExampleMax() {
	fmt.Println(max(-2, 3))
	fmt.Println(max(0, 8))

	// Output:
	// 3
	// 8
}

func ExampleMaxs() {
	fmt.Println(maxs(-2, 3, 0, 1))
	fmt.Println(maxs(8))

	// Output:
	// 3
	// 8
}

func ExampleChmax() {
	ma := 3
	fmt.Println(chmax(&ma, 10))
	fmt.Println(ma)

	ma = 3
	fmt.Println(chmax(&ma, -1))
	fmt.Println(ma)

	// Output:
	// true
	// 10
	// false
	// 3
}
