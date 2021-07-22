package datastruct

import "fmt"

func ExampleSet() {
	// golang does not have 'set' in its standard library.
	// we can use map[T]struct{} as set instead.
	stringSet := make(map[string]struct{})

	// adding elements as the map key
	stringSet["Banana"] = struct{}{}
	stringSet["Apple"] = struct{}{}
	stringSet["Orange"] = struct{}{}

	// check the set contains the "Apple" element
	if _, ok := stringSet["Apple"]; ok {
		fmt.Println("Apple exists")
	}

	// check the set contains the "Berry" element
	if _, ok := stringSet["Berry"]; ok {
		fmt.Println("Berry exists")
	} else {
		fmt.Println("Berry does not exist")
	}

	// run `go test -v` to test this example
	// Output:
	// Apple exists
	// Berry does not exist
}
