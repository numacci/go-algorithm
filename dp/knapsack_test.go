package dp

import "testing"

func TestKnapsackDP(t *testing.T) {
	type args struct {
		n int
		W int
		w []int
		v []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Knapsack DP",
			args: args{
				n: 4,
				W: 5,
				w: []int{2, 1, 3, 2},
				v: []int{3, 2, 4, 2},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KnapsackDP(tt.args.n, tt.args.W, tt.args.w, tt.args.v); got != tt.want {
				t.Errorf("KnapsackDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
