package function

import "fmt"

func ExampleAbs() {
	x := 100
	fmt.Println(Abs(x))

	y := -99
	fmt.Println(Abs(y))

	// Output:
	// 100
	// 99
}
