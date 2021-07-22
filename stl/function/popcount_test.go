package function

import (
	"fmt"
	"math/bits"
)

func ExampleOnesCount() {
	N := 3

	var cnt int
	for bit := 0; bit < (1 << N); bit++ {
		// we can use bits.OnesCount instead of __builtin_popcount in C++
		cnt = bits.OnesCount(uint(bit))
		fmt.Printf("%04b has %d ones\n", bit, cnt)
	}

	// Output:
	// 0000 has 0 ones
	// 0001 has 1 ones
	// 0010 has 1 ones
	// 0011 has 2 ones
	// 0100 has 1 ones
	// 0101 has 2 ones
	// 0110 has 2 ones
	// 0111 has 3 ones
}
