package function

import "fmt"

func ExamplePow() {
	fmt.Println(Pow(4, 3))
	fmt.Println(Pow(-5, 3))
	fmt.Println(Pow(3, 0))

	// Output:
	// 64
	// -125
	// 1
}
