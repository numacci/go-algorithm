package graph

import "fmt"

func ExampleScc() {
	args := []struct {
		from, to int
	}{
		{1, 2},
		{2, 1},
		{2, 3},
		{4, 3},
		{4, 1},
		{1, 4},
		{2, 3},
	}

	n, m := 4, 7
	scc := NewScc(n)
	for i := 0; i < m; i++ {
		a, b := args[i].from, args[i].to
		a--
		b--
		scc.AddEdge(a, b)
	}

	scc.Do()

	mp := make(map[int]int)
	for _, v := range scc.Cmp {
		mp[v]++
	}

	ans := 0
	for _, v := range mp {
		if v == 1 {
			continue
		}
		ans += v * (v - 1) / 2
	}
	fmt.Println(ans)

	// Output:
	// 3
}
