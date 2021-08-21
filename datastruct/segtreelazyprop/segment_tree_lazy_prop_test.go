package segtreelazy

import "fmt"

func ExampleNewSegTreeLazyProp() {
	n := 3

	// RSQ and RAQ
	fx := func(a, b X) X {
		return X{a.val + b.val}
	}
	fm := func(a, b M) M {
		return M{a.val + b.val}
	}
	fa := func(a X, b M) X {
		return X{a.val + b.val}
	}
	fp := func(a M, b int) M {
		return M{a.val * b}
	}
	sg := NewSegTreeLazyProp(n, X{0}, M{0}, fx, fm, fa, fp)

	s, t := 1, 2
	sg.Update(s, t+1, M{1})
	fmt.Println(sg.Query(s, t+1).val)

	// RSQ and RUQ
	fx = func(a, b X) X {
		return X{a.val + b.val}
	}
	fm = func(a, b M) M {
		return b
	}
	fa = func(a X, b M) X {
		return X{b.val}
	}
	fp = func(a M, b int) M {
		return M{a.val * b}
	}
	// set identity element in monoid M to INF in order to distinguish the actual lazy
	sg = NewSegTreeLazyProp(n, X{0}, M{1<<31 - 1}, fx, fm, fa, fp)

	s, t = 2, 5
	sg.Update(s, t+1, M{1})
	s, t = 0, 2
	fmt.Println(sg.Query(s, t+1).val)

	// Output:
	// 2
	// 1
}

func ExampleSegTreeLazyProp() {
	type arg struct {
		com, s, t, x int
	}
	args := []arg{
		{0, 1, 6, -5},
		{0, 2, 4, -9},
		{com: 1, s: 2, t: 3},
		{0, 3, 6, 0},
		{com: 1, s: 0, t: 3},
		{com: 1, s: 5, t: 7},
		{com: 1, s: 2, t: 6},
		{0, 3, 7, 9},
		{com: 1, s: 2, t: 5},
		{0, 0, 1, 1},
	}

	n, q := 8, 10

	fx := func(a, b X) X {
		return X{a.val + b.val}
	}
	fm := func(a, b M) M {
		return b
	}
	fa := func(a X, b M) X {
		return X{b.val}
	}
	fp := func(a M, b int) M {
		return M{a.val * b}
	}
	sg := NewSegTreeLazyProp(n, X{0}, M{1<<31 - 1}, fx, fm, fa, fp)

	for i := 0; i < q; i++ {
		com := args[i].com
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
	// -18
	// -14
	// 0
	// -9
	// 18
}
