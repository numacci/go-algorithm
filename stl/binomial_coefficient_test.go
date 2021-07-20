package stl

import "fmt"

func ExampleBC() {
	// 計算前に初期化が必要
	BCInit()

	fmt.Println(BC(6, 1))
	fmt.Println(BC(10, 2))

	// Output:
	// 6
	// 45
}
