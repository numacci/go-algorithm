package mathematics

// IsPrime is used for primality test
func IsPrime(x int) bool {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Factorize returns the integer factorization result
func Factorize(x int) map[int]int {
	ret := make(map[int]int)
	for i := 2; i*i <= x; i++ {
		for x%i == 0 {
			ret[i]++
			x /= i
		}
	}
	if x != 1 {
		ret[x] = 1
	}
	return ret
}
