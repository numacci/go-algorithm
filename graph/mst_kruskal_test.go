package graph

import "fmt"

func ExampleKruskal() {
	V, E := 7, 9
	es := []Edge{
		{0, 3, 1},
		{1, 2, 10},
		{1, 3, 2},
		{2, 4, 5},
		{3, 4, 7},
		{3, 5, 3},
		{4, 5, 1},
		{4, 6, 8},
		{5, 6, 5},
	}

	cost := Kruskal(V, E, es)
	fmt.Println(cost)

	// Output:
	// 17
}
