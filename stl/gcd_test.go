package stl

import "fmt"

func ExampleGcd() {
	fmt.Println(gcd(6, 9))
	fmt.Println(gcd(5, 7))

	// Output:
	// 3
	// 1
}
