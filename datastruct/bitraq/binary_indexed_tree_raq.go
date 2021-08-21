package bitraq

/*
  Verified:
	RAQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_E/review/5808798/numacci/Go
    RSQ and RAQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_G/review/5808796/numacci/Go
*/

// BIT is 1-indexed binary indexed tree with RAQ
type BIT struct {
	Size int
	Bit  [2][]T // Bit[0]: static add value, Bit[1]: dynamic add value
}

// T used for generics
type T struct {
	val int
}

func NewBIT(n int) *BIT {
	bit := &BIT{
		Size: n + 1,
		Bit:  [2][]T{},
	}
	bit.Bit[0] = make([]T, n+1)
	bit.Bit[1] = make([]T, n+1)
	return bit
}

func (bit *BIT) add(p, k int, x T) {
	for i := k; i < bit.Size; i += i & -i {
		bit.Bit[p][i].val += x.val
	}
}

// Add adds x to [l, r]
func (bit *BIT) Add(l, r int, x T) {
	bit.add(0, l, T{-x.val * (l - 1)})
	bit.add(0, r+1, T{x.val * r})
	bit.add(1, l, T{x.val})
	bit.add(1, r+1, T{-x.val})
}

func (bit *BIT) sum(p, k int) T {
	ret := T{0}
	for i := k; i > 0; i -= i & -i {
		ret.val += bit.Bit[p][i].val
	}
	return ret
}

// Sum returns the sum of [1, k]
func (bit *BIT) Sum(k int) T {
	return T{bit.sum(0, k).val + k*bit.sum(1, k).val}
}

// Query returns the sum of [s, t]
func (bit *BIT) Query(s, t int) T {
	return T{bit.Sum(t).val - bit.Sum(s-1).val}
}
