package stl

import (
	"fmt"
)

// sort.Reverse may not work. Refer to the following site:
//   https://qiita.com/shibukawa/items/0e6e01dc41c352ccedb5
func ExampleReverse() {
	sli := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(sli)/2; i++ {
		sli[i], sli[len(sli)-i-1] = sli[len(sli)-i-1], sli[i]
	}
	fmt.Println(sli)

	// Output:
	// [10 9 8 7 6 5 4 3 2 1]
}
