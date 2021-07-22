package datastruct

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func ExampleSet() {
	// test for int
	intSet := NewSet()
	intSet.Insert(3)
	intSet.Insert(-5)
	intSet.Insert(0)
	fmt.Println(intSet.Contains(-5))
	fmt.Println(intSet.Contains(10))

	intSet.Erase(-5)
	intSet.Insert(10)
	fmt.Println(intSet.Contains(-5))
	fmt.Println(intSet.Contains(10))
	fmt.Println(intSet.Size())

	s := intSet.String()
	s = s[1 : len(s)-1] // remove parentheses
	elements := strings.Fields(s)
	intElems := make([]int, 0, len(elements))
	for _, element := range elements {
		i, _ := strconv.Atoi(element)
		intElems = append(intElems, i)
	}
	sort.Ints(intElems)
	fmt.Println(intElems)

	// test for string
	strSet := NewSet()
	strSet.Insert("banana")
	strSet.Insert("apple")
	strSet.Insert("lemon")
	fmt.Println(strSet.Contains("lemon"))
	fmt.Println(strSet.Contains("peach"))

	strSet.Erase("lemon")
	strSet.Insert("peach")
	fmt.Println(strSet.Contains("lemon"))
	fmt.Println(strSet.Contains("peach"))

	// Output:
	// true
	// false
	// false
	// true
	// 3
	// [0 3 10]
	// true
	// false
	// false
	// true
}
