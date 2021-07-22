package function

import (
	"sort"
)

func LowerBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] >= key })
	default:
		panic("invalid input type for LowerBound")
	}
}

func UpperBound(x interface{}, key interface{}) int {
	switch x.(type) {
	case []int:
		x, key := x.([]int), key.(int)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	case []float64:
		x, key := x.([]float64), key.(float64)
		return sort.Search(len(x), func(i int) bool { return x[i] > key })
	default:
		panic("invalid input type for LowerBound")
	}
}
