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

	G := make([][]*Edge, V)
	for i := 0; i < V; i++ {
		G[i] = make([]*Edge, 0)
	}

	for _, arg := range args {
		G[arg.from] = append(G[arg.from], &Edge{To: arg.to, Cost: arg.cost})
	}

	dist := make([]int, V)
	for i := 0; i < V; i++ {
		dist[i] = 1 << 60
	}
	Dijkstra(G, dist, r)

	for _, v := range dist {
		fmt.Println(v)
	}

	// Output:
	// 0
	// 1
	// 3
	// 4
}
