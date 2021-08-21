package segtree

/*
  Verified:
    RMQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_A/review/5807307/numacci/Go
	RSQ: https://onlinejudge.u-aizu.ac.jp/solutions/problem/DSL_2_B/review/5805828/numacci/Go
*/

// SegTree can be used for RMQ and RSQ with Update operation, not Add.
// So if we need to add x to a[k] when updating the segment tree,
// we need to modify the Update function or use lazy segment tree struct.
type SegTree struct {
	Size int
	Ex   X              // identity element of monoid X: (x,e)=(e,x)=x
	Fx   func(a, b X) X // binary operation of monoid X: (X,X)->X
	Dat  []X
}

// X used for generics and X corresponds to the monoid of this segment tree.
// If we need to construct the float segment tree, change the type of val to float64.
type X struct {
	val int
}

// NewSegTree is a constructor of SegTree. ex is the identity element of monoid X.
// fx is the binary operation of monoid X; e.g. Min, Max, Add.
func NewSegTree(n int, ex X, fx func(X, X) X) *SegTree {
	sg := &SegTree{
		Size: 1,
		Ex:   ex,
		Fx:   fx,
	}

	for sg.Size < n {
		sg.Size *= 2
	}

	// construct segment tree
	sg.Dat = make([]X, 2*sg.Size)
	for i := range sg.Dat {
		sg.Dat[i] = sg.Ex
	}

	return sg
}

// Set and Build used for bulk construction in O(n)
func (sg *SegTree) Set(k int, x X) {
	sg.Dat[k+sg.Size-1] = x
}
func (sg *SegTree) Build() {
	for k := sg.Size - 2; k >= 0; k-- {
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Update updates a[k] to x
func (sg *SegTree) Update(k int, x X) {
	// update itself
	k += sg.Size - 1
	sg.Dat[k] = x

	// propagate to its parent
	for k > 0 {
		k = (k - 1) / 2
		sg.Dat[k] = sg.Fx(sg.Dat[2*k+1], sg.Dat[2*k+2])
	}
}

// Query returns the query result of [s,t) defined by its binary operation Fx
func (sg *SegTree) Query(s, t int) X {
	return sg.query(s, t, 0, 0, sg.Size)
}

func (sg *SegTree) query(s, t, k, l, r int) X {
	if r <= s || t <= l {
		return sg.Ex
	}
	if s <= l && r <= t {
		return sg.Dat[k]
	}

	vl := sg.query(s, t, 2*k+1, l, (l+r)/2)
	vr := sg.query(s, t, 2*k+2, (l+r)/2, r)
	return sg.Fx(vl, vr)
}
