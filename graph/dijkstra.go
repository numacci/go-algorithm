package graph

import (
	"container/heap"
	. "github.com/numacci/go-algorithm/datastruct"
)

// Dijkstra は，ある始点からグラフ上の各点への最短経路を O(E*log(V)) で求める．
// dist は呼び出し元で各要素を INF (1<<60) に初期化する必要があるので注意．
// 優先度付きキューを降順で使う場合にはPへの代入時にマイナスをかければよい．
// Verified:
//   https://onlinejudge.u-aizu.ac.jp/solutions/problem/GRL_1_A/review/5785793/numacci/Go
func Dijkstra(G [][]*Edge, dist []int, s int) {
	h := &PriorityQueue{}
	heap.Init(h)

	dist[s] = 0
	heap.Push(h, &Item{P: 0, V: s})

	for h.Len() > 0 {
		item := heap.Pop(h).(*Item)
		d, v := item.P, item.V.(int)
		if dist[v] < d {
			continue
		}

		for _, nv := range G[v] {
			if dist[nv.To] > dist[v]+nv.Cost {
				dist[nv.To] = dist[v] + nv.Cost
				heap.Push(h, &Item{P: dist[nv.To], V: nv.To})
			}
		}
	}
}
