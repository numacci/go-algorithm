package graph

import (
	. "github.com/numacci/go-algorithm/datastruct"
	"sort"
)

// Kruskal 法によって最小全域木 (Minimum Spanning Tree) を構築する．
// クラスカル法では，予めコストが低い順に辺をソートしておき，
// 辺の数がゼロのグラフにコストの低い辺から追加していく．
// 辺は閉路ができない限り追加し，追加が完了した後次の辺の処理へと進む．
// 閉路の確認にはUnionFindを利用し，各頂点が連結済みかをもって判断する．
// 最小全域木を構築するにあたり，与えられるグラフ es は連結であると仮定する．
func Kruskal(V, E int, es []Edge) int {
	sort.Slice(es, func(i, j int) bool {
		return es[i].Cost < es[j].Cost
	})

	uf := NewUnionFind(V)
	totalCost := 0
	for _, e := range es {
		if !uf.Same(e.From, e.To) {
			uf.Unite(e.From, e.To)
			totalCost += e.Cost
		}
	}
	return totalCost
}
