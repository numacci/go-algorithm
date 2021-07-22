package datastruct

import (
	"fmt"
)

type Set interface {
	Insert(x interface{})
	Erase(x interface{})
	Contains(x interface{}) bool
	Size() int
}

type SetImpl struct {
	data map[interface{}]struct{}
}

func NewSet() *SetImpl {
	return &SetImpl{data: make(map[interface{}]struct{})}
}

func (s *SetImpl) Insert(x interface{}) {
	s.data[x] = struct{}{}
}

func (s *SetImpl) Erase(x interface{}) {
	delete(s.data, x)
}

func (s *SetImpl) Contains(x interface{}) bool {
	_, ok := s.data[x]
	return ok
}

func (s *SetImpl) Size() int {
	return len(s.data)
}

func (s *SetImpl) String() string {
	ret := make([]interface{}, 0, len(s.data))
	for k := range s.data {
		ret = append(ret, k)
	}
	return fmt.Sprint(ret)
}
