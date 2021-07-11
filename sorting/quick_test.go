package sorting

import (
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	type args struct {
		A []int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				A: []int{2, 8, 7, 1, 3, 5, 6, 4},
				p: 0,
				r: 7,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "exercise",
			args: args{
				A: []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11},
				p: 0,
				r: 11,
			},
			want: []int{2, 4, 5, 6, 7, 8, 9, 11, 12, 13, 19, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			QuickSort(tt.args.A, tt.args.p, tt.args.r)
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("QuickSort result = %v, want %v", tt.args.A, tt.want)
			}
		})
	}
}

func TestRandomizedQuickSort(t *testing.T) {
	type args struct {
		A []int
		p int
		r int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				A: []int{2, 8, 7, 1, 3, 5, 6, 4},
				p: 0,
				r: 7,
			},
			want: []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name: "exercise",
			args: args{
				A: []int{13, 19, 9, 5, 12, 8, 7, 4, 21, 2, 6, 11},
				p: 0,
				r: 11,
			},
			want: []int{2, 4, 5, 6, 7, 8, 9, 11, 12, 13, 19, 21},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RandomizedQuickSort(tt.args.A, tt.args.p, tt.args.r)
			if !reflect.DeepEqual(tt.args.A, tt.want) {
				t.Errorf("QuickSort result = %v, want %v", tt.args.A, tt.want)
			}
		})
	}
}
