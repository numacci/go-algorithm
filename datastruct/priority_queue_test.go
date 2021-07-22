package datastruct

import (
	"container/heap"
	"fmt"
)

func ExamplePriorityQueue() {
	items := []*Item{
		{
			v: "banana",
			p: 3,
		},
		{
			v: "apple",
			p: 2,
		},
		{
			v: "pear",
			p: 4,
		},
	}

	// initialize the priority queue
	pq := make(PriorityQueue, 0, len(items))
	for i, item := range items {
		item.index = i
		pq = append(pq, item)
	}
	heap.Init(&pq)

	top := heap.Pop(&pq).(*Item)
	fmt.Println(top.v)

	heap.Push(&pq, &Item{
		v: "orange",
		p: 5,
	})

	for pq.Len() > 0 {
		top = heap.Pop(&pq).(*Item)
		fmt.Println(top.v)
	}

	// Output:
	// pear
	// orange
	// banana
	// apple
}
