package datastruct

import "fmt"

func ExampleUnionFind() {
	const size = 10
	uf := NewUnionFind(size)

	// create {1, 2, 3}
	uf.Unite(1, 2)
	uf.Unite(1, 3)
	// create {4, 5}
	uf.Unite(4, 5)
	// create {7, 8, 9}
	uf.Unite(7, 8)
	uf.Unite(8, 9)

	// union find should have {0}, {1, 2, 3}, {4, 5}, {6}, {7, 8, 9} set at this point
	if uf.Find(1) == uf.Find(2) && uf.Find(1) == uf.Find(3) {
		fmt.Println("1, 2 and 3 belong to the same set")
	}
	if uf.Same(4, 5) {
		fmt.Println("4 and 5 belong to the same set")
	}
	fmt.Printf("size of the set which 7 belongs to: %d\n", uf.Size(7))

	// Output:
	// 1, 2 and 3 belong to the same set
	// 4 and 5 belong to the same set
	// size of the set which 7 belongs to: 3
}
