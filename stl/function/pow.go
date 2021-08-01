package function

func Pow(base, exp int) int {
	if exp == 0 {
		return 1
	}
	res := Pow(base*base, exp/2)
	if exp&1 == 1 {
		res = res * base
	}
	return res
}
