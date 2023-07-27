package go61405_test

import "fmt"

func ExampleRangeOverFunc2() {
	F := func(yield func(id int, name string) bool) bool {
		for _, x := range []struct {
			id   int
			name string
		}{
			{0, "a"},
			{1, "b"},
		} {
			fmt.Println("Before invoke yield", x.id, x.name)
			if result := yield(x.id, x.name); !result {
				return false
			}
			fmt.Println("After invoke yield", x.id, x.name)
		}
		return false
	}
	for k, v := range F {
		fmt.Println(k, v)
	}
	// Output:
	// Before invoke yield 0 a
	// 0 a
	// After invoke yield 0 a
	// Before invoke yield 1 b
	// 1 b
	// After invoke yield 1 b
}
