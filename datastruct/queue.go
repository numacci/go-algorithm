package datastruct

import "fmt"

type Queue interface {
	Push(x interface{})
	Pop() bool
	Front() interface{}
	IsEmpty() bool
}

type QueueImpl struct {
	data []interface{}
	size int
}

func NewQueue(cap int) *QueueImpl {
	return &QueueImpl{data: make([]interface{}, 0, cap), size: 0}
}

func (q *QueueImpl) Push(x interface{}) {
	q.data = append(q.data, x)
	q.size++
}

func (q *QueueImpl) Pop() bool {
	if q.IsEmpty() {
		return false
	}
	q.data[0] = nil
	q.size--
	q.data = q.data[1:]
	return true
}

func (q *QueueImpl) Front() interface{} {
	return q.data[0]
}

func (q *QueueImpl) IsEmpty() bool {
	return q.size == 0
}

func (q *QueueImpl) String() string {
	return fmt.Sprint(q.data)
}
