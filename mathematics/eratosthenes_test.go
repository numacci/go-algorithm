package mathematics

import "fmt"

func ExampleEratosthenes() {
	N := 20
	primes := Eratosthenes(N)
	fmt.Println(primes)

	// Output:
	// [2 3 5 7 11 13 17 19]
}

func ExampleSegmentEratosthenes() {
	a, b := 23, 37
	primes := SegmentEratosthenes(a, b)
	fmt.Println(primes)

	// Output:
	// [23 29 31 37]
}
