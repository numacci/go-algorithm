package function

import "fmt"

func ExampleGcd() {
	fmt.Println(Gcd(6, 9))
	fmt.Println(Gcd(5, 7))

	// Output:
	// 3
	// 1
}
