package mathematics

import "fmt"

func ExampleModPow() {
	fmt.Println(ModPow(2, 4, 4))
	fmt.Println(ModPow(3, 4, 4))

	// Output:
	// 0
	// 1
}
