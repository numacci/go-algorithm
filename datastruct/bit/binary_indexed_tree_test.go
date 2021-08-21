package bit

import "fmt"

func ExampleBIT() {
	type arg struct {
		com, x, y int
	}
	args := []arg{
		{0, 1, 1},
		{0, 2, 2},
		{0, 3, 3},
		{1, 1, 2},
		{1, 2, 2},
	}

	n, q := 3, 5
	b := NewBIT(n)
	for i := 0; i < q; i++ {
		com, x, y := args[i].com, args[i].x, args[i].y
		switch com {
		case 0:
			b.Add(x, T{y})
		case 1:
			fmt.Println(b.Query(x, y).val)
		}
	}

	// Output:
	// 3
	// 2
}
