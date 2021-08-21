package bit

/*
  Verified:
    https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_B/review/5808757/numacci/Go
*/

// BIT is 1-indexed binary indexed tree
type BIT struct {
	Size int
	Bit  []T // 1-indexed
}

// T used for generics
type T struct {
	val int
}

func NewBIT(n int) *BIT {
	return &BIT{
		Size: n + 1,
		Bit:  make([]T, n+1),
	}
}

// Add adds x to k-th element a[k]
func (bit *BIT) Add(k int, x T) {
	for i := k; i < bit.Size; i += i & -i {
		bit.Bit[i].val += x.val
	}
}

// Sum returns the sum of [1, k]
func (bit *BIT) Sum(k int) T {
	ret := T{0}
	for i := k; i > 0; i -= i & -i {
		ret.val += bit.Bit[i].val
	}
	return ret
}

// Query returns the sum of [s, t]
func (bit *BIT) Query(s, t int) T {
	return T{bit.Sum(t).val - bit.Sum(s-1).val}
}
