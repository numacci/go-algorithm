package datastruct

import (
	"container/heap"
	"fmt"
)

func ExamplePriorityQueue() {
	items := []*Item{
		{
			V: "banana",
			P: 3,
		},
		{
			V: "apple",
			P: 2,
		},
		{
			V: "pear",
			P: 4,
		},
	}

	// initialize the priority queue
	pq := make(PriorityQueue, 0, len(items))
	for i, item := range items {
		item.Idx = i
		pq = append(pq, &Item{V: item.V, P: -item.P}) // desc
	}
	heap.Init(&pq)

	top := heap.Pop(&pq).(*Item)
	fmt.Println(top.V)

	heap.Push(&pq, &Item{
		V: "orange",
		P: -5,
	})

	for pq.Len() > 0 {
		top = heap.Pop(&pq).(*Item)
		fmt.Println(top.V)
	}

	// Output:
	// pear
	// orange
	// banana
	// apple
}
