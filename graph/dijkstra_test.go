package graph

import "fmt"

func ExampleDijkstra() {
	V, _, r := 4, 5, 0
	args := []struct {
		from, to, cost int
	}{
		{0, 1, 1},
		{0, 2, 4},
		{1, 2, 2},
		{2, 3, 1},
		{1, 3, 5},
	}

	d := NewDijkstra(V)
	for _, arg := range args {
		d.AddEdge(arg.from, arg.to, arg.cost)
	}

	dist := d.Do(r)

	for _, v := range dist {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 3
	// 4
}
