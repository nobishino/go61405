package collect

import (
	"fmt"

)

func Example() {
	it := iterator(1, 2, 3)
	s := Collect(it)
	for v := range s {
		fmt.Println(v)
	}
	// Output:
	// 0
	// 1
	// 2
}

// 値のみのスライスから、それをyieldするiteratorを返す
func iterator[V any](xs ...V) func(yield func(V) bool) bool {
	return func(yield func(V) bool) bool {
		for _, x := range xs {
			if !yield(x) {
				return false
			}
		}
		return true
	}
}
