package dp

import . "github.com/numacci/go-algorithm/stl"

// KnapsackDP は，以下のような問題を解くときに利用される動的計画法である．
//   重さと価値が定義されたN個の品物の中から重さの総和がWを超えないように
//   いくつか品物を選ぶとき，価値の総和の最大値はいくらになるか．
// 重さの状態を保持しておく必要があるので，DPは以下のように定義する．
//   dp[i+1][j]: i番目までの品物から，重さの総和がjを超えないように
//               品物を選んだ場合の価値の総和の最大値
// 品物を選べなければ価値の総和はゼロなので，DP遷移式の初期条件は
//   dp[0][j] = 0
func KnapsackDP(n, W int, w, v []int) int {
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j <= W; j++ {
			if j < w[i] {
				// i番目の品物がよりj重い→選べない
				dp[i+1][j] = dp[i][j]
			} else {
				// i番目の品物を選べる→選ばない場合と選んだ場合の価値の総和の最大
				dp[i+1][j] = Max(dp[i][j], dp[i][j-w[i]]+v[i])
			}
		}
	}
	return dp[n][W]
}
