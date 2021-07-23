package mathematics

import (
	. "github.com/numacci/go-algorithm/stl/function"
	"math"
)

// Eratosthenes returns prime numbers in [1, n]
func Eratosthenes(n int) []int {
	primes := make([]int, 0, n)

	isPrime := make([]bool, n+1)
	for i := 2; i <= n; i++ {
		isPrime[i] = true
	}

	for i := 2; i <= n; i++ {
		if !isPrime[i] {
			continue
		}

		primes = append(primes, i)
		for j := 2 * i; j <= n; j += i {
			// exclude multiples of prime number 'i'
			isPrime[j] = false
		}
	}
	return primes
}

// SegmentEratosthenes returns prime numbers in [a, b]
func SegmentEratosthenes(a, b int) []int {
	isPrimeAtoB := make([]bool, b-a+1)                           // sieve in [a, b]
	isPrimeToSqrtB := make([]bool, int(math.Sqrt(float64(b)))+1) // sieve in [0, sqrt(b))

	for i := 0; i*i <= b; i++ {
		isPrimeToSqrtB[i] = true
	}
	for i := 0; i <= b-a; i++ {
		isPrimeAtoB[i] = true
	}

	for i := 2; i*i <= b; i++ {
		if !isPrimeToSqrtB[i] {
			continue
		}
		for j := 2 * i; j*j <= b; j += i {
			isPrimeToSqrtB[j] = false
		}
		for j := Max(i, (a+i-1)/i) * i; j <= b; j += i {
			isPrimeAtoB[j-a] = false
		}
	}

	primes := make([]int, 0, b-a+1)
	for i := a; i <= b; i++ {
		if isPrimeAtoB[i-a] {
			primes = append(primes, i)
		}
	}
	return primes
}
