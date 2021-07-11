package stl

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxs(x ...int) int {
	if len(x) <= 0 {
		return 0
	}
	ret := x[0]
	for _, v := range x {
		if ret < v {
			ret = v
		}
	}
	return ret
}

func chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
