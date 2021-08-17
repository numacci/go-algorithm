package graph

import (
	"container/heap"
	. "github.com/numacci/go-algorithm/datastruct"
)

/*
  Verified:
    https://onlinejudge.u-aizu.ac.jp/solutions/problem/GRL_1_A/review/5793328/numacci/Go
*/

// Dijkstra は，始点sから全頂点への最短経路を O(E*log(V)) で求めるアルゴリズム．
// 優先度付きキューを降順で使う場合にはPへの代入時にマイナスをかければよい．
type Dijkstra struct {
	Inf  int
	G    [][]*Edge
	Dist []int
}

func NewDijkstra(n int) *Dijkstra {
	inf := 1 << 60

	dist := make([]int, n)
	for i := 0; i < n; i++ {
		dist[i] = inf
	}

	return &Dijkstra{
		Inf:  inf,
		G:    make([][]*Edge, n),
		Dist: dist,
	}
}

func (d *Dijkstra) AddEdge(from, to, cost int) {
	d.G[from] = append(d.G[from], &Edge{To: to, Cost: cost})
}

func (d *Dijkstra) Do(s int) []int {
	h := &PriorityQueue{}
	heap.Init(h)

	d.Dist[s] = 0
	heap.Push(h, &Item{P: 0, V: s})

	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		v := item.V.(int)
		if d.Dist[v] < item.P {
			continue
		}

		for _, nv := range d.G[v] {
			if d.Dist[nv.To] > d.Dist[v]+nv.Cost {
				d.Dist[nv.To] = d.Dist[v] + nv.Cost
				heap.Push(h, &Item{P: d.Dist[nv.To], V: nv.To})
			}
		}
	}
	return d.Dist
}
