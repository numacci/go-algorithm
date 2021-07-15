package dp

// DigitDP は，以下のような問題を解くときに利用される動的計画法である．
//   「0以上N以下の整数で，ある条件を満たす整数の個数を求めよ」
//   「0以上N以下の整数で，ある条件を満たす最大の整数を求めよ」
//   ref. https://qiita.com/pinokions009/items/1e98252718eeeeb5c9ab
// ここでは上記記事に従い，0以上N以下の整数で，いずれかの桁に3を含む整数の個数を求める．
func DigitDP(N string) int {
	// digits: Nの各桁の数を保持
	digits := make([]int, len(N))
	for i := 0; i < len(N); i++ {
		digits[i] = int(N[i] - '0')
	}

	// dp[i][smaller][j]: 下記条件を満たす整数の総数
	//   i:       今調べている桁 (上から)
	//   smaller: 未満フラグ (今まで調べてきた数の並びがN未満だと確定している＝1)
	//   j:       今まで調べてきた数の中に3が含まれる＝1
	// 最終的な答えは dp[len(N)][0][1] + dp[len(N)][1][1]
	dp := make([][2][2]int, len(N)+1)

	// 初期条件
	dp[0][0][0] = 1

	// 上の位から桁を固定していき，条件を満たす整数の総数を計算する
	for i := 0; i < len(N); i++ {
		for smaller := 0; smaller < 2; smaller++ {
			for j := 0; j < 2; j++ {
				// i桁目の数字として動かせる範囲は，
				// 未満フラグが正であれば0から9までの全ての数，
				// そうでなければNのi桁目の数
				d := 9
				if smaller == 0 {
					d = digits[i]
				}
				// i桁目の数を動かす
				for x := 0; x <= d; x++ {
					// i桁目を考えるとき，i-1桁目までに未満フラグが正であれば正
					// そうでない場合，動かしている数xがNのi桁目より小さければ正
					//   e.g. 12345の3桁目のdpを計算しているときに，x=2 (=122**)
					nextSmaller := 0
					if smaller == 1 || x < digits[i] {
						nextSmaller = 1
					}
					// 今までに3が出ている or xが3であれば次のjは正
					nextJ := 0
					if j == 1 || x == 3 {
						nextJ = 1
					}
					// 状態遷移先はnextSmallerとnextJで計算される
					// dp[i][smaller][j]はその遷移先の状態数として加算される
					dp[i+1][nextSmaller][nextJ] += dp[i][smaller][j]
				}
			}
		}
	}
	return dp[len(N)][0][1] + dp[len(N)][1][1]
}
