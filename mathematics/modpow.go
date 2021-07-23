package mathematics

func ModPow(base, exp, mod int) int {
	ret := 1
	for exp > 0 {
		if exp&1 == 1 {
			ret = ret * base % mod
		}
		base = base * base % mod
		exp >>= 1
	}
	return ret
}
