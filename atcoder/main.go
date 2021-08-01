package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer w.Flush()

	// Implement from here
}

// IO Utils
var (
	sc = bufio.NewScanner(os.Stdin)
	w  = bufio.NewWriter(os.Stdout)
)

const (
	initBufSize = 1024 * 1024
	maxBufSize  = 1024 * 1024 * 1024
)

func init() {
	buf := make([]byte, initBufSize)
	sc.Buffer(buf, maxBufSize)
	sc.Split(bufio.ScanWords)
}

func ns() string {
	sc.Scan()
	return sc.Text()
}

func ni() int {
	x, _ := strconv.Atoi(ns())
	return x
}

func nf() float64 {
	x, _ := strconv.ParseFloat(ns(), 64)
	return x
}

func nss(n int) []string {
	ret := make([]string, n)
	for i := 0; i < n; i++ {
		ret[i] = ns()
	}
	return ret
}

func nis(n int) []int {
	ret := make([]int, n)
	for i := 0; i < n; i++ {
		ret[i] = ni()
	}
	return ret
}

func nfs(n int) []float64 {
	ret := make([]float64, n)
	for i := 0; i < n; i++ {
		ret[i] = nf()
	}
	return ret
}

func out(x ...interface{}) {
	fmt.Fprintln(w, x...)
}

func Abs(x int) int {
	if x < 0 {
		x *= -1
	}
	return x
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a / Gcd(a, b) * b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Maxs(x ...int) int {
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

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Mins(x ...int) int {
	if len(x) <= 0 {
		return 0
	}
	ret := x[0]
	for _, v := range x {
		if ret > v {
			ret = v
		}
	}
	return ret
}

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

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
