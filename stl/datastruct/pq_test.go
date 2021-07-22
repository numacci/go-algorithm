package datastruct

import (
	"container/heap"
	"fmt"
)

func ExampleIntHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)

	fmt.Println(heap.Pop(h).(int))

	heap.Push(h, 3)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h).(int))
	}

	// Output:
	// 5
	// 3
	// 2
	// 1
}
