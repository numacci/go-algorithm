package segtree

import (
	"fmt"
)

func ExampleNewSegTree() {
	const inf = 1 << 60

	// length of the sequence
	n := 3

	// Ranged Minimum Query
	min := func(a, b X) X {
		if a.val < b.val {
			return a
		}
		return b
	}
	rmq := NewSegTree(n, X{inf}, min)
	rmq.Update(1, X{2})
	rmq.Query(0, 2)

	// Ranged Maximum Query
	max := func(a, b X) X {
		if a.val > b.val {
			return a
		}
		return b
	}
	rmq = NewSegTree(n, X{-inf}, max)
	rmq.Update(1, X{2})
	rmq.Query(0, 2)

	// Ranged Sum Query
	sum := func(a, b X) X {
		return X{a.val + b.val}
	}
	rsq := NewSegTree(n, X{0}, sum)
	rsq.Update(1, X{2})
	rsq.Query(0, 2)
}

// Example for Ranged Minimum Query
func ExampleSegTree() {
	type arg struct {
		com, x, y int
	}

	n, q := 3, 5
	args := []arg{
		{0, 0, 1},
		{0, 1, 2},
		{0, 2, 3},
		{1, 0, 2},
		{1, 1, 2},
	}

	const ex = 1<<31 - 1
	min := func(a, b X) X {
		if a.val < b.val {
			return a
		}
		return b
	}
	sg := NewSegTree(n, X{ex}, min)
	for i := 0; i < q; i++ {
		com, x, y := args[i].com, args[i].x, args[i].y
		switch com {
		case 0:
			sg.Update(x, X{y})
		case 1:
			ans := sg.Query(x, y+1)
			fmt.Println(ans.val)
		}
	}

	// Output:
	// 1
	// 2
}
