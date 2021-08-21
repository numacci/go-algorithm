package segtreelazy

import "fmt"

func ExampleNewSegTreeLazy() {
	n := 3
	// RMQ and RUQ
	fx := func(a, b X) X {
		if a.val < b.val {
			return a
		}
		return b
	}
	fm := func(a, b M) M {
		return b
	}
	fa := func(a X, b M) X {
		return X{b.val}
	}
	sg := NewSegTreeLazy(n, X{1<<31 - 1}, M{1<<31 - 1}, fx, fm, fa)

	sg.Update(0, 1, M{1})
	fmt.Println(sg.Query(0, 2).val)

	// RMQ and RAQ
	n = 6

	fx = func(a, b X) X {
		if a.val < b.val {
			return a
		}
		return b
	}
	fm = func(a, b M) M {
		return M{a.val + b.val}
	}
	fa = func(a X, b M) X {
		return X{a.val + b.val}
	}
	sg = NewSegTreeLazy(n, X{1<<31 - 1}, M{0}, fx, fm, fa)

	// initialize segment tree element by 0, while setting the identity element of monoid M to INF
	for i := 0; i < n; i++ {
		sg.Set(i, X{0})
	}
	sg.Build()

	sg.Update(1, 3, M{1})
	sg.Update(2, 4, M{-2})
	fmt.Println(sg.Query(0, 5).val)

	// Output:
	// 1
	// -2
}

func ExampleSegTreeLazy() {
	type arg struct {
		op, s, t, x int
	}
	args := []arg{
		{0, 0, 1, 1},
		{0, 1, 2, 3},
		{0, 2, 2, 2},
		{op: 1, s: 0, t: 2},
		{op: 1, s: 1, t: 2},
	}

	n, q := 3, 5

	fx := func(a, b X) X {
		if a.val < b.val {
			return a
		}
		return b
	}
	fm := func(a, b M) M {
		return b
	}
	fa := func(a X, b M) X {
		return X{b.val}
	}
	sg := NewSegTreeLazy(n, X{1<<31 - 1}, M{1<<31 - 1}, fx, fm, fa)

	for i := 0; i < q; i++ {
		com := args[i].op
		switch com {
		case 0:
			s, t, x := args[i].s, args[i].t, args[i].x
			sg.Update(s, t+1, M{x})
		case 1:
			s, t := args[i].s, args[i].t
			fmt.Println(sg.Query(s, t+1).val)
		}
	}

	// Output:
	// 1
	// 2
}
