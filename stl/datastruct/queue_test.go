package datastruct

import "fmt"

func ExampleQueue() {
	// we can use slice instead of queue
	s := make([]int, 0)

	// Push
	s = append(s, 1)
	s = append(s, 8)
	s = append(s, -2)

	// Pop
	front := s[0]
	s = s[1:]
	fmt.Println(front)

	front = s[0]
	s = s[1:]
	fmt.Println(front)

	s = append(s, 10)
	front = s[0]
	s = s[1:]
	fmt.Println(front)

	front = s[0]
	s = s[1:]
	fmt.Println(front)

	// Output:
	// 1
	// 8
	// -2
	// 10
}
