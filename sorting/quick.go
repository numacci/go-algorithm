package sorting

import (
	"math/rand"
	"time"
)

func QuickSort(A []int, p, r int) {
	if p < r {
		q := partition(A, p, r)
		QuickSort(A, p, q-1)
		QuickSort(A, q+1, r)
	}
}

func partition(A []int, p, r int) int {
	x := A[r]  // pivot
	i := p - 1 // pivotより小さい要素の集合の上端
	for j := p; j < r; j++ {
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}
	// pivotより大きい要素の集合と小さい要素の集合の間にpivotを配置
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

func RandomizedQuickSort(A []int, p, r int) {
	if p < r {
		q := randomizedPartition(A, p, r)
		RandomizedQuickSort(A, p, q-1)
		RandomizedQuickSort(A, q+1, r)
	}
}

func randomizedPartition(A []int, p, r int) int {
	// change pivot randomly
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(r-p+1) + p
	A[i], A[r] = A[r], A[i]

	return partition(A, p, r)
}
