package stl

import "fmt"

func ExampleStack() {
	s := make([]int, 0)

	// Push
	s = append(s, 1)
	s = append(s, 8)
	s = append(s, -2)

	// Pop
	top := s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(top)

	top = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(top)

	s = append(s, 10)
	top = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(top)

	top = s[len(s)-1]
	s = s[:len(s)-1]
	fmt.Println(top)

	// Output:
	// -2
	// 8
	// 10
	// 1
}
