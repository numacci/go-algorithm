package mathematics

import "fmt"

func ExampleIsPrime() {
	for i := 2; i <= 10; i++ {
		if IsPrime(i) {
			fmt.Printf("%d is prime!\n", i)
		} else {
			fmt.Printf("%d is not prime...\n", i)
		}
	}

	// Output:
	// 2 is prime!
	// 3 is prime!
	// 4 is not prime...
	// 5 is prime!
	// 6 is not prime...
	// 7 is prime!
	// 8 is not prime...
	// 9 is not prime...
	// 10 is not prime...
}

func ExampleFactorize() {
	N := 24
	for k, v := range Factorize(N) {
		fmt.Printf("%d has %d^%d\n", N, k, v)
	}

	N = 37
	for k, v := range Factorize(N) {
		fmt.Printf("%d has %d^%d\n", N, k, v)
	}

	// Output:
	// 24 has 2^3
	// 24 has 3^1
	// 37 has 37^1
}
