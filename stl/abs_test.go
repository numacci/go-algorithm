package stl

import "fmt"

func ExampleAbs() {
	x := 100
	fmt.Println(abs(x))

	y := -99
	fmt.Println(abs(y))

	// Output:
	// 100
	// 99
}
