package function

import "fmt"

func ExampleNextPermutation() {
	items := []int{1, 1, 3, 4}
	for {
		fmt.Println(items)
		if !NextPermutation(items) {
			break
		}
	}

	// Output:
	// [1 1 3 4]
	// [1 1 4 3]
	// [1 3 1 4]
	// [1 3 4 1]
	// [1 4 1 3]
	// [1 4 3 1]
	// [3 1 1 4]
	// [3 1 4 1]
	// [3 4 1 1]
	// [4 1 1 3]
	// [4 1 3 1]
	// [4 3 1 1]
}
