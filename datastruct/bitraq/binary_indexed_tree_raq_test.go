package bitraq

import "fmt"

func ExampleBIT() {
	type arg struct {
		com, s, t, x int
	}
	args := []arg{
		{0, 1, 2, 1},
		{0, 2, 3, 2},
		{0, 3, 3, 3},
		{com: 1, s: 1, t: 2},
		{com: 1, s: 2, t: 3},
	}

	n, q := 3, 5
	bit := NewBIT(n)
	for i := 0; i < q; i++ {
		com := args[i].com
		switch com {
		case 0:
			s, t, x := args[i].s, args[i].t, args[i].x
			bit.Add(s, t, T{x})
		case 1:
			s, t := args[i].s, args[i].t
			fmt.Println(bit.Query(s, t).val)
		}
	}

	// Output:
	// 4
	// 8
}
