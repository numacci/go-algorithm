package datastruct

import "fmt"

type Stack interface {
	Push(x interface{})
	Pop() bool
	Top() interface{}
	IsEmpty() bool
}

type StackImpl struct {
	data []interface{}
	size int
}

func NewStack(cap int) *StackImpl {
	return &StackImpl{data: make([]interface{}, 0, cap), size: 0}
}

func (s *StackImpl) Push(x interface{}) {
	s.data = append(s.data, x)
	s.size++
}

func (s *StackImpl) Pop() bool {
	if s.IsEmpty() {
		return false
	}
	s.data[s.size-1] = nil
	s.size--
	s.data = s.data[:s.size]
	return true
}

func (s *StackImpl) Top() interface{} {
	return s.data[s.size-1]
}

func (s *StackImpl) IsEmpty() bool {
	return s.size == 0
}

func (s *StackImpl) String() string {
	return fmt.Sprint(s.data)
}
