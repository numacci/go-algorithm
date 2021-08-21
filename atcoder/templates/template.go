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

/* ============================
           TEMPLATE
============================ */

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

// Standard Libraries
