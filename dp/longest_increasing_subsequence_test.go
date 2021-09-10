package dp

import "testing"

func TestLisDP(t *testing.T) {
	type args struct {
		n int
		a []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "p64 case",
			args: args{
				n: 5,
				a: []int{4, 2, 3, 1, 5},
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LisDP(tt.args.n, tt.args.a); got != tt.want {
				t.Errorf("LisDP() = %v, want %v", got, tt.want)
			}
		})
	}
}
