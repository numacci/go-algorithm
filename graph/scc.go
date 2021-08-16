package graph

/*
  Verified:
    https://atcoder.jp/contests/typical90/submissions/25090679
*/

type Scc struct {
	N       int     // グラフの頂点数
	G       [][]int // グラフの隣接リスト表現
	RG      [][]int // 辺の向きを逆にしたグラフ
	Vs      []int   // 帰りがけ順の並び
	Visited []bool
	Cmp     []int // グラフに属する強連結成分のトポロジカル順序
}

func NewScc(n int) *Scc {
	return &Scc{
		N:       n,
		G:       make([][]int, n),
		RG:      make([][]int, n),
		Vs:      make([]int, 0),
		Visited: make([]bool, n),
		Cmp:     make([]int, n),
	}
}

func (s *Scc) AddEdge(from, to int) {
	s.G[from] = append(s.G[from], to)
	s.RG[to] = append(s.RG[to], from)
}

func (s *Scc) Dfs(v int) {
	s.Visited[v] = true
	for _, nv := range s.G[v] {
		if !s.Visited[nv] {
			s.Dfs(nv)
		}
	}
	s.Vs = append(s.Vs, v) // 帰りがけ順で頂点番号をメモ
}

func (s *Scc) Rdfs(v, k int) {
	s.Visited[v] = true
	s.Cmp[v] = k
	for _, pv := range s.RG[v] {
		if !s.Visited[pv] {
			s.Rdfs(pv, k)
		}
	}
}

func (s *Scc) Do() int {
	// 1回目のDFS
	s.Visited = make([]bool, s.N)
	for v := 0; v < s.N; v++ {
		if !s.Visited[v] {
			s.Dfs(v)
		}
	}

	// 1回目と逆順でDFS
	s.Visited = make([]bool, s.N)
	k := 0 // 強連結成分をまとめてDAGにしたときの頂点数
	for i := s.N - 1; i >= 0; i-- {
		if !s.Visited[s.Vs[i]] {
			// Rdfsでは強連結成分が同じkを持つように更新される
			s.Rdfs(s.Vs[i], k)
			k++
		}
	}
	return k
}
