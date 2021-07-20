package stl

const (
	BCMax = 510000
	BCMod = 1000000007
)

var fac, finv, inv [BCMax]int

// BCInit initializes the binomial coefficient table
func BCInit() {
	fac[0], fac[1] = 1, 1
	finv[0], finv[1] = 1, 1
	inv[1] = 1
	for i := 2; i < BCMax; i++ {
		fac[i] = fac[i-1] * i % BCMod
		inv[i] = BCMod - inv[BCMod%i]*(BCMod/i)%BCMod
		finv[i] = finv[i-1] * inv[i] % BCMod
	}
}

// BC calculates the binomial coefficient nCr
func BC(n, r int) int {
	if n < r {
		return 0
	}
	if n < 0 || r < 0 {
		return 0
	}
	// n!/(r!*(n-r)!) = n!*(r!)^{-1}*((n-r)!)^{-1}
	return fac[n] * (finv[r] * finv[n-r] % BCMod) % BCMod
}
