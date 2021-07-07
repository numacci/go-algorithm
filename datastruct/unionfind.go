package datastruct

import "fmt"

type UnionFind struct {
	par  []int
	rank []int
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		par:  make([]int, n),
		rank: make([]int, n),
	}
	for i := 0; i < n; i++ {
		uf.par[i] = -1
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	if uf.par[x] < 0 {
		return x
	} else {
		uf.par[x] = uf.Find(uf.par[x])
		return uf.par[x]
	}
}

func (uf *UnionFind) Unite(x, y int) {
	x = uf.Find(x)
	y = uf.Find(y)
	if x == y {
		return
	}

	if uf.rank[x] < uf.rank[y] {
		x, y = y, x
	}
	if uf.rank[x] == uf.rank[y] {
		uf.rank[x]++
	}

	uf.par[x] += uf.par[y]
	uf.par[y] = x
}

func (uf *UnionFind) Same(x, y int) bool {
	return uf.Find(x) == uf.Find(y)
}

func (uf *UnionFind) Size(x int) int {
	return -uf.par[uf.Find(x)]
}

func (uf *UnionFind) String() string {
	return fmt.Sprintf("par = %v, rank = %v", uf.par, uf.rank)
}
