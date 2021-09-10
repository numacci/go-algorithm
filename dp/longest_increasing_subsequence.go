package dp

import (
	"math"
	"sort"
)

// LisDP case 1 O(N^2)
//func LisDP(n int, a []int) int {
//	max := 1
//	dp := make([]int, n)
//	for i := range dp {
//		dp[i] = math.MaxInt32
//	}
//
//	for i := 0; i < n; i++ {
//		dp[i] = 1
//		for j := 0; j < i; j++ {
//			if a[j] < a[i] {
//				dp[i] = Max(dp[i], dp[j]+1)
//			}
//			Chmax(&max, dp[i])
//		}
//	}
//	return max
//}

// LisDP case 2 O(N*logN)
func LisDP(n int, a []int) int {
	// 長さがi+1であるような増加部分列における最終要素の最小値
	dp := make([]int, n)
	for i := range dp {
		dp[i] = math.MaxInt32
	}

	for i := 0; i < n; i++ {
		idx := sort.Search(n, func(j int) bool {
			return a[i] < dp[j]
		})
		dp[idx] = a[i]
	}
	return sort.Search(n, func(i int) bool {
		return dp[i] == math.MaxInt32
	})
}
