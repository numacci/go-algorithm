package function

func NextPermutation(arr []int) bool {
	// calculate the left index to swap
	l := len(arr) - 2
	for l >= 0 && arr[l] >= arr[l+1] {
		l--
	}
	if l < 0 {
		return false
	}
	// calculate the right index to swap
	r := len(arr) - 1
	for arr[l] >= arr[r] {
		r--
	}
	// swap
	arr[l], arr[r] = arr[r], arr[l]
	// reverse elements between l and r
	for l, r = l+1, len(arr)-1; l < r; l, r = l+1, r-1 {
		arr[l], arr[r] = arr[r], arr[l]
	}
	return true
}
