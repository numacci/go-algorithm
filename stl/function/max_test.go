package function

import "fmt"

func ExampleMax() {
	fmt.Println(Max(-2, 3))
	fmt.Println(Max(0, 8))

	// Output:
	// 3
	// 8
}

func ExampleMaxs() {
	fmt.Println(Maxs(-2, 3, 0, 1))
	fmt.Println(Maxs(8))

	// Output:
	// 3
	// 8
}

func ExampleChmax() {
	ma := 3
	fmt.Println(Chmax(&ma, 10))
	fmt.Println(ma)

	ma = 3
	fmt.Println(Chmax(&ma, -1))
	fmt.Println(ma)

	// Output:
	// true
	// 10
	// false
	// 3
}
