package datastruct

import "fmt"

func ExampleStack() {
	// test for type int
	intStack := NewStack(10)

	intStack.Push(10)
	intStack.Push(1)
	intStack.Push(-5)

	fmt.Println(intStack.Top())
	intStack.Pop()

	intStack.Push(5)
	for !intStack.IsEmpty() {
		fmt.Println(intStack.Top())
		intStack.Pop()
	}

	// test for type struct
	type pair struct {
		first, second int
	}
	structStack := NewStack(10)

	structStack.Push(&pair{first: 3, second: 10})
	structStack.Push(&pair{first: -1, second: 2})
	structStack.Push(&pair{first: 5, second: -5})

	top := structStack.Top().(*pair)
	fmt.Println(top.first, top.second)
	structStack.Pop()

	structStack.Push(&pair{0, 3})
	for !structStack.IsEmpty() {
		top = structStack.Top().(*pair)
		fmt.Println(top.first, top.second)
		structStack.Pop()
	}

	// Output:
	// -5
	// 5
	// 1
	// 10
	// 5 -5
	// 0 3
	// -1 2
	// 3 10
}
