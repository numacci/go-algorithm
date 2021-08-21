package segtreelazy

/*
  Verified:
	RUQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_D/review/5807775/numacci/Go
	RAQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_E/review/5807895/numacci/Go
    RMQ and RUQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_F/review/5807767/numacci/Go
	RMQ and RAQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_H/review/5808026/numacci/Go
*/

type SegTreeLazy struct {
	Size int
	Ex   X  // identity element in monoid X
	Em   M  // identity element in monoid M
	Fx   FX // binary operation in monoid X
	Fm   FM // binary operation in monoid M
	Fa   FA // binary operation between X and M
	Dat  []X
	Lazy []M
}

type X struct {
	val int
}
type M struct {
	val int
}
type FX func(a, b X) X
type FM func(a, b M) M
type FA func(a X, b M) X

func NewSegTreeLazy(n int, ex X, em M, fx FX, fm FM, fa FA) *SegTreeLazy {
	sg := &SegTreeLazy{
		Size: 1,
		Fx:   fx,
		Fm:   fm,
		Fa:   fa,
		Ex:   ex,
		Em:   em,
	}

	for sg.Size < n {
		sg.Size *= 2
	}
	sg.Dat = make([]X, 2*sg.Size)
	for i := range sg.Dat {
		sg.Dat[i] = sg.Ex
	}
	sg.Lazy = make([]M, 2*sg.Size)
	for i := range sg.Lazy {
		sg.Lazy[i] = sg.Em
	}
	return sg
}

// Set and Build used for bulk construction in O(n)
func (sg *SegTreeLazy) Set(k int, x X) {
	sg.Dat[k+sg.Size-1] = x
}
func (sg *SegTreeLazy) Build() {
	for k := sg.Size - 2; k >= 0; k-- {
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Eval propagates the lazy data
func (sg *SegTreeLazy) Eval(k int) {
	if sg.Lazy[k] == sg.Em {
		return
	}

	// propagate k-th lazy to its children
	if k < sg.Size-1 {
		sg.Lazy[2*k+1] = sg.Fm(sg.Lazy[2*k+1], sg.Lazy[k])
		sg.Lazy[2*k+2] = sg.Fm(sg.Lazy[2*k+2], sg.Lazy[k])
	}
	// update itself
	sg.Dat[k] = sg.Fa(sg.Dat[k], sg.Lazy[k])
	sg.Lazy[k] = sg.Em
}

func (sg *SegTreeLazy) update(a, b int, x M, k, l, r int) {
	sg.Eval(k)
	if a <= l && r <= b {
		sg.Lazy[k] = sg.Fm(sg.Lazy[k], x)
		sg.Eval(k)
	} else if a < r && l < b {
		sg.update(a, b, x, 2*k+1, l, (l+r)/2)
		sg.update(a, b, x, 2*k+2, (l+r)/2, r)
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Update updates [a, b) to x
func (sg *SegTreeLazy) Update(a, b int, x M) {
	sg.update(a, b, x, 0, 0, sg.Size)
}

func (sg *SegTreeLazy) query(a, b, k, l, r int) X {
	sg.Eval(k)
	if r <= a || b <= l {
		return sg.Ex
	}
	if a <= l && r <= b {
		return sg.Dat[k]
	}
	vl := sg.query(a, b, 2*k+1, l, (l+r)/2)
	vr := sg.query(a, b, 2*k+2, (l+r)/2, r)
	return sg.Fx(vl, vr)
}

// Query returns the query result in [a, b)
func (sg *SegTreeLazy) Query(a, b int) X {
	return sg.query(a, b, 0, 0, sg.Size)
}
