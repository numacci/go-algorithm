package mathematics

import (
	"fmt"
	"sort"
)

func ExampleDivisor() {
	N := 24

	div := Divisor(N)
	sort.Ints(div)
	fmt.Println(div)

	// Output:
	// [1 2 3 4 6 8 12 24]
}
