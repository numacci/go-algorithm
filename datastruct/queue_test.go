package datastruct

import "fmt"

func ExampleQueue() {
	// test for type int
	intQueue := NewQueue(10)

	intQueue.Push(10)
	intQueue.Push(1)
	intQueue.Push(-5)

	fmt.Println(intQueue.Front())
	intQueue.Pop()

	intQueue.Push(5)
	for !intQueue.IsEmpty() {
		fmt.Println(intQueue.Front())
		intQueue.Pop()
	}

	// test for type struct
	type pair struct {
		first, second int
	}
	structQueue := NewQueue(10)

	structQueue.Push(&pair{first: 3, second: 10})
	structQueue.Push(&pair{first: -1, second: 2})
	structQueue.Push(&pair{first: 5, second: -5})

	top := structQueue.Front().(*pair)
	fmt.Println(top.first, top.second)
	structQueue.Pop()

	structQueue.Push(&pair{0, 3})
	for !structQueue.IsEmpty() {
		top = structQueue.Front().(*pair)
		fmt.Println(top.first, top.second)
		structQueue.Pop()
	}

	// Output:
	// 10
	// 1
	// -5
	// 5
	// 3 10
	// -1 2
	// 5 -5
	// 0 3
}
