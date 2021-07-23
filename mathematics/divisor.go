package mathematics

func Divisor(x int) []int {
	ret := make([]int, 0, x)
	for i := 1; i*i <= x; i++ {
		if x%i == 0 {
			ret = append(ret, i)
			if i != x/i {
				ret = append(ret, x/i)
			}
		}
	}
	return ret
}
